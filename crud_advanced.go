package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func crudAdvanced(db *gorm.DB) {
	fmt.Println("============================")
	fmt.Println("CRUD Advanced")
	fmt.Println("============================")

	// Create
	user := User{
		FirstName: "Arthur",
		LastName:  "Dent",
		Username:  "adent",
		Salary:    5000,
	}

	// Check if it is a new record
	fmt.Println(db.NewRecord(&user))
	db.Create(&user)

	// Find record that does not exist
	// Use custom function to see if it extsts
	u := User{}
	db.Find(&u, 100)
	fmt.Println(u.NotFound())

	// Query multiple records
	users := []User{}
	db.Where(&User{Salary: 200}).Find(&users)
	fmt.Println(users)

	// Update with column names
	db.Model(&user).Update("first_name", "zipp")
	db.Model(&user).Updates(
		map[string]interface{}{
			"first_name": "Zap",
			"last_name":  "Bee",
		})
	// Will not trigger callback
	db.Model(&user).UpdateColumn("first_name", "zipp")
	db.Model(&user).UpdateColumns(
		map[string]interface{}{
			"first_name": "Zap",
			"last_name":  "Smith",
		})

	// Batch Update
	db.Table("users").Where("last_name = ?", "Smith").
		Update("first_name", "Penny")

	// Update based on values in the columns
	db.Table("users").Where("salary > ?", 3000).
		Update("salary", gorm.Expr("salary + 500"))

		// Select one record and delete it
	db.Table("users").Where("salary > ?", 3000).Delete(&User{})
	// Select all records from a model and delete all
	db.Model(&User{}).Delete(&User{})

	// Delete record using ordinary deletetion
	db.Unscoped().Delete(&user)

	// Using Transaction to create
	tx := db.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
	}

	// Using Transaction to update
	user.LastName = "Happy Robot"
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
	}

	// Run sql once this function is called
	tx.Commit()
}
