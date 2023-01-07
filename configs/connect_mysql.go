package configs

import (
	"database/sql"
	"log"
)

func GetDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@/fiber")
	// Run docker vs container
	// db, err := sql.Open("mysql", "root:root@tcp(database:3306)/fiber")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
