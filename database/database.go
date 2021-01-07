package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ChiKangMa/go-support/debug"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

var Connection, DataSourceName string

func InitDB() {
	var err error
	driverName, dataSourceName := getDBConfig()
	Db, err = sql.Open(driverName, dataSourceName)
	debug.PrintError(err)
	if err = Db.Ping(); err != nil {
		log.Panic(err)
	}
}

func CloseDB() {
	Db.Close()
}

func getDBConfig() (string, string) {
	if hasInitialized() == false {
		return loadDBConfig()
	}
	return Connection, DataSourceName
}

func hasInitialized() bool {
	return len(Connection) > 0 && len(DataSourceName) > 0
}

func loadDBConfig() (string, string) {
	Connection = os.Getenv("DB_CONNECTION")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	DataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		username,
		password,
		host,
		port,
		database)
	return Connection, DataSourceName
}
