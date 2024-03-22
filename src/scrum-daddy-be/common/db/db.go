package db

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

type Database struct {
	db *gorm.DB
}

func (db *Database) GetDB() *gorm.DB {
	return db.db
}

func Connect() *Database {
	connectionString := CreateConnectionString()
	dbConnection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		slog.Error("Error connecting to database: ", "error", err)
	}
	return &Database{db: dbConnection}
}

func (db *Database) Close() {
	sqlDB, err := db.db.DB()
	if err != nil {
		slog.Error(err.Error())
	}
	err = sqlDB.Close()
	if err != nil {
		slog.Error(err.Error())
	}
}

type DatabaseConnection struct {
	Db *gorm.DB
}

func (db *DatabaseConnection) WithContext(ctx context.Context) *gorm.DB {
	return db.Db.WithContext(ctx)
}

func (db *DatabaseConnection) Get() *gorm.DB {
	return db.Db
}

type UnitOfWork struct {
	db *gorm.DB
}

type IUnitOfWork interface {
	Commit() error
	Rollback() error
	GetConnection() *DatabaseConnection
}

func (uow *UnitOfWork) GetConnection() *DatabaseConnection {
	return &DatabaseConnection{Db: uow.db}
}

func (uow *UnitOfWork) Commit() error {
	if uow.db == nil {
		slog.Error("No transaction to commit")
		return nil
	}
	err := uow.db.Commit().Error
	uow.db = nil
	return err
}

func (uow *UnitOfWork) Rollback() error {
	if uow.db == nil {
		slog.Error("No transaction to rollback")
		return nil
	}
	err := uow.db.Rollback().Error
	uow.db = nil
	return err
}

type UowFactoryFn func(db *Database) (IUnitOfWork, error)

// UowTransactionalFactory creates a new unit of work and starts a new transaction
func UowTransactionalFactory(db *Database) (IUnitOfWork, error) {
	tx := db.GetDB().Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &UnitOfWork{db: tx}, nil
}

func UowFactory(db *Database) (IUnitOfWork, error) {
	dbInstance := db.GetDB()
	if dbInstance.Error != nil {
		return nil, dbInstance.Error
	}
	return &UnitOfWork{db: dbInstance}, nil
}
