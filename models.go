package main

import (
// "github.com/jinzhu/gorm"
)

type User struct {
	ID uint

	// Set column type
	Username  string `sql:"type:VARCHAR(255)"`
	FirstName string

	// Set default value
	LastName string `sql:"DEFAULT:'Smith'"`

	// Ignored attribute will be treated as attr instead of column
	IgnoredField bool `sql:"-"`

	// Relationship
	Calendar Calendar

	// Others

	// auto-populate columns id and timestamps
	// gorm.Model
	// Model gorm.Model `gorm:"embedded"`

	// Custom primary ket
	// UserID int `gorm:"primary_key"`

	// Custom column name
	// FirstName string `gorm:"column:FirstName"`

	// AUTO_INCREMENT can only be set on key field
	// Count     int `gorm:"AUTO_INCREMENT"`

	// Not Null & Unique field
	// UniqueField string `sql:"not null;unique"`
}

type Calendar struct {
	ID     uint
	Name   string
	UserID uint
}

var users []User = []User{
	User{Username: "foobar", FirstName: "Foo", LastName: "Bar"},
	User{Username: "helloworld", FirstName: "Hello", LastName: "World"},
	User{Username: "john", FirstName: "John"},
}
