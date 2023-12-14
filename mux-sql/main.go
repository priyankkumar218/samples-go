package main

import (
	"log"
)

func main() {
	a := &App{}
	err := a.Initialize(
		"mux-sql-postgres-1", // postgres host
		//Change to `mux-sql-postgres-1` when using Docker to run keploy
		"postgres", // username
		"password", // password
		"postgres") // db_name

	if err != nil {
		log.Fatal("Failed to initialize app", err)
	}

	log.Printf("😃 Connected to 8010 port !!")

	a.Run(":8010")
}
