package conf

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"time"
)

func InitMysqlDB() (*sql.DB, error) {
	var (
		errMysql error
		dbConn   *sql.DB
	)

	dbHost := os.Getenv("HOST")
	dbPort := os.Getenv("PORT")
	dbUser := os.Getenv("USER_NAME")
	dbPass := os.Getenv("PASS")
	dbName := os.Getenv("NAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, errMysql = sql.Open(`mysql`, dsn)

	dbConn.SetMaxOpenConns(300)
	dbConn.SetMaxIdleConns(25)
	dbConn.SetConnMaxLifetime(5 * time.Minute)

	return dbConn, errMysql
}
