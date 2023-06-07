package conf

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"
)

func InitMysqlDB() (*sql.DB, error) {
	var (
		errMysql error
		dbConn   *sql.DB
	)

	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "kolaborasisalt_kolaborasisalt"
	dbPass := "Ky4F-E*Yb^XT"
	dbName := "kolaborasisalt_donor_darah"

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
