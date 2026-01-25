package database

import (
	//mysql is for mysql connection driver
	"time"

	"github.com/iamgenii/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DataStore defines database connection
// here we are using gorm for connection
type DataStore struct {
	db *gorm.DB
}

// NewDataStore for create mysql database connection
// get and connection string as parameter
// client connection and return a connection
func NewDataStore(connectionStr string, dbConfig configs.DBConfig) (*gorm.DB, error) {

	//Open connection using gorm
	client, err := gorm.Open("mysql", connectionStr)
	if err != nil {
		panic("Error in establish database client connection")
	}
	client.LogMode(dbConfig.DBLogMode)
	client.DB().SetMaxIdleConns(dbConfig.DbConnectionPool.MaxIdealConnection)
	client.DB().SetConnMaxLifetime(time.Duration(dbConfig.DbConnectionPool.MaxConnectionLifetimeInMinutes) * time.Minute)
	client.DB().SetMaxOpenConns(dbConfig.DbConnectionPool.MaxOpenConnection)
	return client, nil

}
