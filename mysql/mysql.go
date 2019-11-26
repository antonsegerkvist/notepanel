package mysql

import (
	"database/sql"
	"log"
	"strconv"
)

const username = "webuser"
const password = "dummypassword"
const host = "127.0.0.1"
const port = 3306
const dbname = "notepanel"

var database *sql.DB

func mysqlToFormatDNS() string {
	return username + ":" + password +
		"@tcp(" + host + ":" + strconv.FormatInt(int64(port), 10) +
		")/" + dbname
}

func init() {
	var err error
	database, err = sql.Open("mysql", mysqlToFormatDNS())
	if err != nil {
		log.Fatal("==> Error: " + err.Error())
	}

	database.SetMaxOpenConns(20)
	database.SetMaxIdleConns(20)
}

//
// Open gets a database handler to the mysql server.
//
func Open() (*sql.DB, error) {

	err := database.Ping()
	if err != nil {
		err = database.Ping()
		if err != nil {
			err = database.Ping()
		}
	}
	return database, err

}
