package main

import (
	"meepshop_project/database"
	"meepshop_project/service/server"
	"meepshop_project/utils/config"
)

func main() {
	config.InitConfig()
	database.InitDB()
	server.RunServer()
}
