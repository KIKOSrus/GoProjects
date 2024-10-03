package main

import (
	"TaskManager/cmd"
	"TaskManager/db"
	"log"
)

func main() {
	dbPath := "tasks.db"

	err := db.Init(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	cmd.Execute()

}
