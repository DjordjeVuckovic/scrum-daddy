package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	DB *gorm.DB
}

func (db *Database) GetDB() *gorm.DB {
	return db.DB
}

func Connect() *Database {
	connectionString := CreateConnectionString()
	dbConnection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return &Database{DB: dbConnection}
}

func (db *Database) Close() {
	sqlDB, err := db.DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	err = sqlDB.Close()
	if err != nil {
		log.Fatal(err)
	}
}

type UnitOfWork interface {
	BeginTransaction() *Database
	Rollback() error
}

func (db *Database) BeginTransaction() *Database {
	return &Database{DB: db.DB.Begin()}
}

func Rollback(db *Database) error {
	err := db.DB.Rollback().Error
	if err != nil {
		log.Fatal(err)
	}
	return err
}
