package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xvbnm48/echo-restfull/config"
)

var db *sql.DB

var err error

func Init() {
	conf := config.GetConfig()

	connection := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err = sql.Open("mysql", connection)
	if err != nil {
		panic("Error connecting to database")
	}

	err = db.Ping()
	if err != nil {
		panic("Error pinging database")
	}
}

func createCon() *sql.DB {
	return db
}
