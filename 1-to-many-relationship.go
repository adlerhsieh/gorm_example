package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func oneToManyRelationship(db *gorm.DB) {
	fmt.Println("============================")
	fmt.Println("1-to-many Relationship")
	fmt.Println("============================")

	db.CreateTable(&Appointment{})

	db.Save(&User{
		FirstName: "Blake",
		LastName:  "William",
		Username:  "wblake",
		Calendar: Calendar{
			Name: "Another Calendar",
			// Created only if there is a foreign key tag specified
			Appointments: []Appointment{
				{Subject: "foo", Description: "bar"},
				{Subject: "baz", Description: "qux"},
			},
		},
	})

	u := User{FirstName: "Blake"}
	c := Calendar{}
	a := []Appointment{}
	// Queries have to split
	db.Find(&u).Related(&c, "calendar")
	db.Find(&c).Related(&a, "appointments")
	fmt.Println("User -> Calendar -> Appointment")
	fmt.Println(a)
}
