package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func columnNames(db *gorm.DB) {
	fmt.Println("============================")
	fmt.Println("Column Names")
	fmt.Println("============================")

	for _, f := range db.NewScope(&User{}).Fields() {
		println(f.Name)
	}
}
