package main

import (
	"log"

	"github.com/3b3al/Charity-System/database"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// Database Connection
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("connect to database: %v\n", err)
	}

	// Migrate Step
	
	err = database.Migrate(db)
	if err != nil {
		log.Fatalf("migrate database: %v\n", err)
	}
	
}
