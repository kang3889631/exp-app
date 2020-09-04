package gdo

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"

	"srv/core/config"
)

var (
	// DB ...
	DB *sqlx.DB
)

// Init ...
func Init() error {
	var err error

	DB, err = sqlx.Connect(config.Data.DB.Driver, config.Data.DB.DataSourceName)

	if err != nil {
		log.Printf("Are we inside here? :%s", err.Error())
		return err
	}

	// TODO: move setting to config
	DB.SetMaxOpenConns(1000)
	DB.SetMaxIdleConns(300)
	DB.SetConnMaxLifetime(0)

	return err
}

// SetSQLWaitTimeout ...
func SetSQLWaitTimeout(db *sqlx.DB, timeout int) (err error) {
	if _, err = db.Exec("SET @old_wait_timeout := @@session.wait_timeout"); err != nil {
		return err
	}

	if _, err = db.Exec("SET @@session.wait_timeout := ?", timeout); err != nil {
		return err
	}

	return nil
}

// ResetSQLWaitTimeout ...
func ResetSQLWaitTimeout(db *sqlx.DB) (err error) {
	if _, err = db.Exec("SET @@session.wait_timeout := @old_wait_timeout"); err != nil {
		return err
	}

	return nil
}

// Transaction ...
// NOTE: it can also use named return variable for error.
//
// Example:
// err = gdo.Transaction(func(tx *sqlx.Tx) error {
// 	var err error
//
// 	if _, err = tx.NamedExec("UPDATE tb SET Changed = NOW() WHERE ID = :ID", map[string]interface{}{"ID": 1}); err != nil {
// 		return err
// 	}
//
// 	return nil
// })
func Transaction(txFunc func(*sqlx.Tx) error, timeout int) error {
	return transaction(DB, txFunc, timeout)
}

// TransactionDB ...
func TransactionDB(db *sqlx.DB, txFunc func(*sqlx.Tx) error, timeout int) error {
	return transaction(db, txFunc, timeout)
}

func transaction(db *sqlx.DB, txFunc func(*sqlx.Tx) error, timeout int) (err error) {
	if err = SetSQLWaitTimeout(db, timeout); err != nil {
		return err
	}

	//
	var tx *sqlx.Tx

	if tx, err = db.Beginx(); err != nil {
		return err
	}

	defer func(tx *sqlx.Tx, e *error) {
		if tx == nil {
			*e = fmt.Errorf("tx is nil")
			return
		}

		if p := recover(); p != nil {
			// Set *e ONLY when robackErr has an error.
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				*e = rollbackErr
				return
			}

			// re-throw panic after Rollback
			panic(p)
		} else if *e != nil {
			// Set *e ONLY when robackErr has an error.
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				*e = rollbackErr
				return
			}
		} else {
			// err is nil; if Commit returns error update err
			if *e = tx.Commit(); *e != nil {
				return
			}
		}

		// NOTE: Ignoring the error
		_ = ResetSQLWaitTimeout(db)
	}(tx, &err)

	return txFunc(tx)
}
