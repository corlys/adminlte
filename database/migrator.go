package database

import (
	"fmt"

	"github.com/corlys/adminlte/app/entity"
	"github.com/corlys/adminlte/database/seeder"
	"gorm.io/gorm"
)

func DBMigrate(db *gorm.DB) {
	err := db.AutoMigrate(entity.User{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if err := DBSeed(db); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func DBSeed(db *gorm.DB) error {
	err := seeder.Userseeder(db)
	if err != nil {
		return err
	}
	return nil
}
