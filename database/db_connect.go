package database

import (
	"gin_bank/utils/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

const (
	GormQueryOption = "gorm:query_option"
	SyntaxForUpdate = "FOR UPDATE"
)

var Db *gorm.DB
var (
	GormSetSelectForUpdate = func() (string, string) { return GormQueryOption, SyntaxForUpdate }
)

func InitDB() {
	var err error
	dataSourceName := config.GlobalConfig.GetDatabaseUrl()
	Db, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
}

// NewTestSession only for unit-test
func NewTestSession() *gorm.DB {
	return Db.New()
}
