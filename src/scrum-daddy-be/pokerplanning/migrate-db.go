package pokerplanning

import (
	"log"
	"scrum-daddy-be/common/db"
	"scrum-daddy-be/pokerplanning/types"
)

func MigrateDb(db *db.Database) {
	err := db.GetDB().AutoMigrate(&types.PokerRoom{})
	if err != nil {
		log.Fatal(err)
	}
}
