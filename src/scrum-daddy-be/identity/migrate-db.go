package identity

import (
	"log"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/identity/domain"
)

func MigrateDb(db *db.Database) {
	err := db.GetDB().AutoMigrate(&domain.User{}, &domain.Role{})
	if err != nil {
		log.Fatal(err)
	}
}
