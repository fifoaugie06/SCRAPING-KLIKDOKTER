package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root@/simcovid?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}

	return db
}
