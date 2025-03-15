package database

import (
	"log"

	"gorm.io/gorm/logger"

	"github.com/glebarez/sqlite"
	"github.com/mosescode1/BASIC_GO_CRUD.git/internals/entity"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dbName string) (*gorm.DB, error){
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}


func CloseConnection() {
	if DB == nil {
		return
	}

	dbInterface, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = dbInterface.Close()
	if err != nil {
		log.Fatal(err)
	}
}


func MigrateDtabase() error {
	if err := DB.AutoMigrate(&entity.Todo{}); err != nil {
		return err
	}
	return nil
}