package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ColumnNames() {
	fmt.Println("============================")
	fmt.Println("Column Names")
	fmt.Println("============================")
	db, err := gorm.Open("mysql", "gorm:gorm@/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for _, f := range db.NewScope(&User{}).Fields() {
		println(f.Name)
	}
}
