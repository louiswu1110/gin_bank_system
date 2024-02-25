package main

import (
	"meepshop_project/database"
	"meepshop_project/service/server"
)

func main() {
	database.InitDB()
	server.RunServer()
}
