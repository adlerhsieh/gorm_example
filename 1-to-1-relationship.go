package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func oneToOneRelationship(db *gorm.DB) {
	fmt.Println("============================")
	fmt.Println("1-to-1 Relationship")
	fmt.Println("============================")

	db.CreateTable(&Calendar{})

	// Add constraint foreign key
	// Don't have to do this if we're only querying
	db.Model(&Calendar{}).
		AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	// Create record and related record
	db.Save(&User{
		FirstName: "Jane",
		LastName:  "Eyre",
		Username:  "jane",
		Calendar: Calendar{
			Name: "Base",
		},
	})

	// This user has no related calendar
	u := User{}
	c := Calendar{}
	db.First(&u).Related(&c, "calendar")
	fmt.Println(c)

	// This user has related calendar
	u = User{}
	c = Calendar{}
	db.Last(&u).Related(&c, "calendar")
	fmt.Println(c)
}
