package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func Open() *sql.DB {
	connStr := "user=buddy_admin password=mysecretpassword dbname=tg_bot_database sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("connected")
		return db
	}
	return nil
}

func Close(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Panic("connection is not close")
	}
}
