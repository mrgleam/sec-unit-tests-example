package main

import (
	"github.com/mrgleam/sec-unit-tests-example/database"
	"github.com/mrgleam/sec-unit-tests-example/server"
)

func main() {
	db := database.SetupDB()
	server.Server(db)
}
