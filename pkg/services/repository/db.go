package repository

import (
	"log"

	"github.com/mccune1224/job-tracker/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDatabase(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func MigrateDatabase(db *gorm.DB) {
	for _, m := range models.RegisterModels() {
		err := db.Debug().AutoMigrate(m.Model)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	log.Print("Successfully Migrated Models")
}
