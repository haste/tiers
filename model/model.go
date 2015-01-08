package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"tiers/conf"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", conf.Config.Database)
	if err != nil {
		log.Fatalf("Error on initializing database connection: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error on opening database connection: %s", err.Error())
	}
}
