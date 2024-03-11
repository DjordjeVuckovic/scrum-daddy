package common

import (
	"scrum-daddy-be/common/db"
)

func MigrateDb(db *db.Database) {
	//err := db.GetDB().AutoMigrate(&identity.User{}, &identity.Role{})
	//if err != nil {
	//	log.Fatal(err)
	//}
}
