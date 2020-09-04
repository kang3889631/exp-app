package OnlineOrder

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"srv/core/gdo"
	"srv/core/util"
)

// SrvPostOnlineOrder ...
func SrvPostOnlineOrder(c *gin.Context) {
	// Delcare err variable
	var err error

	// Prepare a response data
	data := map[string]interface{}{
		"status": false,
	}

	// Insert user submitted data into database
	// TODO: Need more implementaiton
	if err = InsertOnlineOrder(); err != nil {
		// If err, return err
		data["error"] = err.Error()
		c.SecureJSON(http.StatusOK, data)

		return
	}

	// If ok, set status to true
	data["status"] = true

	// Output data as a JSON string back to the client
	c.SecureJSON(http.StatusOK, data)

	return
}

// SrvGetOnlineOrder ...
func SrvGetOnlineOrder(c *gin.Context) {
	// Delcare err variable
	var err error

	// Prepare a response data
	data := map[string]interface{}{
		"status": false,
	}

	// Get product data from database
	var prodArr interface{}

	if prodArr, err = QryOnlineOrder(); err != nil {
		// If err, return err
		data["error"] = err.Error()
		c.SecureJSON(http.StatusOK, data)

		return
	}

	// Set to response data
	data["productArr"] = prodArr

	// If ok, set status to true
	data["status"] = true

	// Output data as a JSON string back to the client
	c.SecureJSON(http.StatusOK, data)

	return
}

// QryOnlineOrder ...
func QryOnlineOrder() (interface{}, error) {
	var err error
	var ns *sqlx.NamedStmt
	var rows *sqlx.Rows

	type ColDef struct {
		Name  string  `db:"Name"`
		Price float32 `db:"Price"`
	}

	qry := `
	SELECT
	Name
	, Price
	FROM Product
	WHERE
	ProductID > :ProductID
	`

	if ns, err = gdo.DB.PrepareNamed(qry); err != nil {
		return nil, err
	}

	defer util.DeferClose(ns, &err)

	//
	if rows, err = ns.Queryx(map[string]interface{}{
		"ProductID": "0",
	}); err != nil {
		return nil, err
	}

	defer util.DeferClose(rows, &err)

	//
	colDef := ColDef{}
	rowData := []ColDef{}

	for rows.Next() {
		if err = rows.StructScan(&colDef); err != nil {
			return nil, err
		}

		rowData = append(rowData, colDef)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rowData, nil
}

// InsertOnlineOrder ...
func InsertOnlineOrder() error {
	var err error
	var ns *sqlx.NamedStmt
	//var rs sql.Result

	//
	qry := `
	INSERT INTO OnlineOrder (CustomerID, Amount, Tax, ShippingCost, TotalAmount, Created)
	VALUES (:CustomerID, :Amount, :Tax, :ShippingCost, :TotalAmount, NOW())
	`

	// Prepared MySQL code
	if ns, err = gdo.DB.PrepareNamed(qry); err != nil {
		return err
	}

	defer util.DeferClose(ns, &err)

	//
	if _, err = ns.Exec(map[string]interface{}{
		"CustomerID":   2,
		"Amount":       0,
		"Tax":          0,
		"ShippingCost": 0,
		"TotalAmount":  0,
	}); err != nil {
		return err
	}

	//rs.LastInsertId()
	//rs.RowsAffected()

	return nil
}
