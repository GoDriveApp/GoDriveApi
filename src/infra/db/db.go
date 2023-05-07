package db

import (
	"github.com/GoDriveApp/GoDriveApi/src/core/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbURL = "postgres://postgres:postgrespw@localhost:32768/GoDriveDb"
)

func Init() (error, *gorm.DB) {

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return err, nil
	}

	err = db.AutoMigrate(&entity.Account{}, &entity.User{})
	if err != nil {
		return err, nil
	}

	return nil, db
}
