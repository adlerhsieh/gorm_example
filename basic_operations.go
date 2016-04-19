package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func BasicOperations() {
	fmt.Println("============================")
	fmt.Println("Basic Operations")
	fmt.Println("============================")
	// Connection
	db, err := gorm.Open("mysql", "gorm:gorm@/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	dbase := db.DB()
	defer dbase.Close()

	// Ensure Connection
	err = dbase.Ping()
	if err != nil {
		panic(err.Error())
	}

	// Create & Drop table
	if db.HasTable(&User{}) {
		db.DropTable(&User{})
	}
	db.CreateTable(&User{})

	// Seeding db using custom model data
	for _, user := range users {
		db.Create(&user)
	}

	// Select first & last user in table
	firstUser := User{}
	db.First(&firstUser)
	fmt.Println("First User:")
	fmt.Println(firstUser)

	lastUser := User{}
	db.Last(&lastUser)
	fmt.Println("Last User:")
	fmt.Println(lastUser)

	// Where & Update
	u := User{FirstName: "Foo"}
	// Where 用設定的條件下去查，First會把結果綁到變數上
	db.Where(&u).First(&u)
	fmt.Println("User before update:")
	fmt.Println(u)

	u.LastName = "Beeblebrox"
	db.Save(&u)

	q := User{}
	db.Where(&u).First(&q)
	fmt.Println("User after update:")
	fmt.Println(q)

	// Delete
	d := User{Username: "foobar"}
	db.Where(&d).Delete(&User{})
	db.Where(&d).First(&d)
	fmt.Println("Looking for deleted user:")
	// 用已知的查詢條件，去資料庫查再綁回同一個struct上
	// 查不到的話就return原本的查詢條件
	fmt.Println(d)
	fmt.Println("============================")
}

func (u User) TableName() string {
	// custom table name, this is default
	return "users"
}
