package main

import (
	"meepshop_project/database"
	"meepshop_project/service/server"
	"meepshop_project/utils/config"
	"meepshop_project/utils/tools"
)

func main() {
	tools.InitGenerator()
	config.InitConfig()
	database.InitDB()
	server.RunServer()
}
