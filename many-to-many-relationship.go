package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func manyToManyRelationship(db *gorm.DB) {
	fmt.Println("============================")
	fmt.Println("many-to-many Relationship")
	fmt.Println("============================")

	var users []User
	db.Where("Username in (?)", []string{"john", "jane"}).Find(&users)

	db.Save(&User{
		FirstName: "Johnny",
		LastName:  "Depp",
		Username:  "jdepp",
		Calendar: Calendar{
			Name: "B Calendar",
			Appointments: []Appointment{
				{Subject: "beats", Description: "audio", Attendees: users},
				{Subject: "Apple", Description: "iPhone", Attendees: users},
			},
		},
	})

	// Query associated users from Appointment
	a := Appointment{Subject: "beats"}
	u := []User{}
	db.Where(&a).First(&a)
	// use column names for querying association
	db.Model(&a).Association("Attendees").Find(&u)
	fmt.Println(u)

	// Query associated appointments from User
	usr := User{Username: "john"}
	apps := []Appointment{}
	db.Where(&usr).First(&usr)
	db.Model(&usr).Association("Appointments").Find(&apps)
	fmt.Println(apps)

}
