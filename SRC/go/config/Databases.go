package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func DatabaseInit() {
	var err error

	db, err = sql.Open("postgres", "user=theodelaune dbname=goapi")

	if err != nil {
		log.Fatal(err)
	}

	// Create Table pays if not exists
	createPaysTable()
}

func createPaysTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS pays(id serial,names varchar(20), capital varchar(20), monnaie varchar(20), Languages int, Région timestamp default NULL, SousRégion timestamp default NULL, constraint pk primary key(id))")

	if err != nil {
		log.Fatal(err)
	}
}

// Getter for db var
func Db() *sql.DB {
	return db
}
