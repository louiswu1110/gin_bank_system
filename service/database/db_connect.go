package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"meepshop_project/utils/config"
)

var db *gorm.DB

func InitDB() {
	var err error
	dataSourceName := config.GlobalConfig.GetDatabaseUrl()
	db, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
}
