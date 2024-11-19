package repository

import (
	"database/sql"
	"log"
	"walking-information-storage-server-api/database"
)

var db = ConnectDB()

func ConnectDB() *sql.DB {
	db, err := database.SqlConnect()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
