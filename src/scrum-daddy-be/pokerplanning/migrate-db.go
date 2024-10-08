package pokerplanning

import (
	"log"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/pokerplanning/domain"
)

func MigrateDb(db *db.Database) {
	err := db.GetDB().AutoMigrate(&domain.PokerRoom{})
	if err != nil {
		log.Fatal(err)
	}
}
