package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := Connect()
	purgeDB(db)
	crud(db)
	columnNames(db)
	oneToOneRelationship(db)
	oneToManyRelationship(db)
	manyToManyRelationship(db)
	defer db.Close()
}

func purgeDB(db *gorm.DB) {
	if db.HasTable(&Calendar{}) {
		db.DropTable(&Calendar{})
	}
	if db.HasTable(&User{}) {
		db.DropTable(&User{})
	}
	if db.HasTable(&Appointment{}) {
		db.DropTable(&Appointment{})
	}
}
