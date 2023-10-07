package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dbConn *sql.DB

func SetupConnection() {
	conn, err := sql.Open("mysql", "root:@/url_shortener")
	if err != nil {
		panic(err)
	}

	dbConn = conn
	dbConn.SetConnMaxLifetime(time.Minute * 3)
	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(10)
}

func DBConn() *sql.DB {
	return dbConn
}
