package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	// Data ...
	Data config
)

type config struct {
	Sys struct {
		Mode         string
		Version      string
		CasbinModel  string
		CasbinPolicy string
	}
	DB struct {
		Driver         string
		DataSourceName string
	}
	Web struct {
		Domain    string
		BindAddr  string
		HTTPSPort string
		HTTPPort  string
	}
}

// Init ...
func Init() error {
	var err error

	// Try to read the config from the following locations
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	if err = viper.ReadInConfig(); err != nil {
		return err
	}

	if err = viper.Unmarshal(&Data); err != nil {
		return err
	}

	// TODO: ensure config watch function works as expected
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		if e.Op == fsnotify.Write {
			log.Printf("Config file changed (%s): %s\n", e.Op, e.Name)

			// Try to decode first
			tmpData := config{}

			if err := viper.Unmarshal(&tmpData); err != nil {
				log.Fatal(err)
			}

			if !tmpData.IsSysModeValid() {
				log.Fatal(fmt.Errorf("Sys.Mode %s is invalid\n", Data.Sys.Mode))
			}

			// Decode
			if err := viper.Unmarshal(&Data); err != nil {
				log.Fatal(err)
			}
		}
	})

	if !Data.IsSysModeValid() {
		return fmt.Errorf("Sys.Mode %s is invalid\n", Data.Sys.Mode)
	}

	return nil
}

// IsSysModeValid ...
func (c config) IsSysModeValid() bool {
	if c.Sys.Mode == "dev" || c.Sys.Mode == "prod" {
		return true
	}

	return false
}

// IsSysModeDev ...
func (c config) IsSysModeDev() bool {
	return c.Sys.Mode == "dev"
}

// WebDomainHTTPSPort ...
func (c config) WebDomainHTTPSPort() string {
	return fmt.Sprintf("%s:%s", c.Web.Domain, c.Web.HTTPSPort)
}

// WebDomainHTTPPort ...
func (c config) WebDomainHTTPPort() string {
	return fmt.Sprintf("%s:%s", c.Web.Domain, c.Web.HTTPPort)
}

// WebBindAddrHTTPSPort ...
func (c config) WebBindAddrHTTPSPort() string {
	return fmt.Sprintf("%s:%s", c.Web.BindAddr, c.Web.HTTPSPort)
}

// WebBindAddrHTTPPort ...
func (c config) WebBindAddrHTTPPort() string {
	return fmt.Sprintf("%s:%s", c.Web.BindAddr, c.Web.HTTPPort)
}
