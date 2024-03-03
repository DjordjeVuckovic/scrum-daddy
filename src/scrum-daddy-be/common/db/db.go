package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	DB *gorm.DB
}

func (g *Database) GetDB() *gorm.DB {
	return g.DB
}

func Connect() *Database {
	connectionString := CreateConnectionString()
	dbConnection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return &Database{DB: dbConnection}
}

func Close(db *Database) {
	sqlDB, err := db.DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	err = sqlDB.Close()
	if err != nil {
		log.Fatal(err)
	}
}
