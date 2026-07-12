package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() {

	var err error

	DB, err = sql.Open(
		"sqlite3",
		"korzadivpn.db",
	)

	if err != nil {

		log.Fatal(err)

	}

	err = DB.Ping()

	if err != nil {

		log.Fatal(err)

	}

	log.Println("Base de datos conectada")

}
