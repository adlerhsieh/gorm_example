package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func crud(db *gorm.DB) {
	fmt.Println("============================")
	fmt.Println("CRUD")
	fmt.Println("============================")

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

	// Add index with index_name, column_name
	db.Model(&User{}).AddIndex("idx_first_name", "first_name")
	db.Model(&User{}).AddUniqueIndex("idx_username", "username")
	// Remove index name
	db.Model(&User{}).RemoveIndex("idx_first_name")
}

func (u User) TableName() string {
	// custom table name, this is default
	return "users"
}
