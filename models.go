package main

import (
// "github.com/jinzhu/gorm"
)

type User struct {
	ID        uint
	Username  string `sql:"type:VARCHAR(255)"`
	FirstName string
	// Set default value
	LastName string `sql:"DEFAULT:'Smith'"`
	// ignored attribute will be treated as attr instead of column
	IgnoredField bool `sql:"-"`
	// Not Null & Unique field
	UniqueField string `sql:"not null;unique"`

	// Others

	// auto-populate columns id and timestamps
	// gorm.Model
	// Custom primary ket
	// UserID int `gorm:"primary_key"`
	// Custom column name
	// FirstName string `gorm:"column:FirstName"`
	// AUTO_INCREMENT can only be set on key field
	// Count     int `gorm:"AUTO_INCREMENT"`
}

var users []User = []User{
	User{Username: "foobar", FirstName: "Foo", LastName: "Bar", UniqueField: "a"},
	User{Username: "helloworld", FirstName: "Hello", LastName: "World", UniqueField: "b"},
	User{Username: "john", FirstName: "John", UniqueField: "c"},
}
