package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() {

	var err error
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "korzadivpn.db"
	}

	DB, err = sql.Open(
		"sqlite3",
		dbPath,
	)

	if err != nil {

		log.Fatal(err)

	}

	err = DB.Ping()

	if err != nil {

		log.Fatal(err)

	}

	log.Printf("Base de datos conectada en: %s", dbPath)

}
