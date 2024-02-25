package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"meepshop_project/utils/config"
)

var Db *gorm.DB

func InitDB() {
	var err error
	dataSourceName := config.GlobalConfig.GetDatabaseUrl()
	Db, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
}
