package controllers

import "database/sql"

var DB *sql.DB

func InitDB(db *sql.DB) {
	DB = db
}
