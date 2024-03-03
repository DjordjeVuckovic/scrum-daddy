package pokerplanning

import (
	"log"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/pokerplanning/rooms"
)

func MigrateDb(db *db.Database) {
	err := db.GetDB().AutoMigrate(&rooms.PokerRoom{})
	if err != nil {
		log.Fatal(err)
	}
}
