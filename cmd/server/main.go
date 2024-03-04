package main

import (
	"gin_bank/database"
	"gin_bank/service/server"
	"gin_bank/utils/config"
	"gin_bank/utils/tools"
)

func main() {
	tools.InitGenerator()
	config.InitConfig()
	database.InitDB()
	server.RunServer()
}
