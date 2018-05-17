package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
)

type Configuration struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	GRPCPort   string
	TimeZone   string
}

func CreateDBConnection(c Configuration) (*gorm.DB, error) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", c.TimeZone)
	dbConnectionString := fmt.Sprintf("%s?%s", connectionString, val.Encode())
	return gorm.Open(
		"mysql",
		dbConnectionString,
	)
}
