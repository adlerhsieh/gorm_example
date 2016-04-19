package main

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string
	FirstName string
	LastName  string
}

var users []User = []User{
	User{Username: "foobar", FirstName: "Foo", LastName: "Bar"},
	User{Username: "helloworld", FirstName: "Hello", LastName: "World"},
	User{Username: "john", FirstName: "John", LastName: "Smith"},
}
