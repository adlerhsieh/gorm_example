package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(
		"mysql",
		"gorm:gorm@/gorm?charset=utf8&parseTime=True&loc=Local",
	)
	if err != nil {
		panic(err.Error())
	}
	return db
}
