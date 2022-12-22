package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

var db *sqlx.DB

func MustInitialize(dbPath string) {
	var err error

	db, err = sqlx.Open("sqlite", dbPath)

	if err != nil {
		log.Fatalf("Failed to open the database connection: %s", err.Error())
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("Failed to pint the database: %s", err.Error())
	}

}

func Transaction(cb func(tx *sqlx.Tx) error) error {
	tx, err := db.Beginx()

	if err != nil {
		return err
	}
	err = cb(tx)

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func Close() {
	db.Close()
}
