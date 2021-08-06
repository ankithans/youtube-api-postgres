package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ankithans/youtube-api/pkg/models"
)

var (
	DBConn *gorm.DB
)

func NewDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "postgres://user:password@db:5432/db?sslmode=disable",
	}), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	db.AutoMigrate(models.Video{})

	return db
}
