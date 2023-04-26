package database

import (
	"log"
	"os"

	"github.com/grahamkatana/fiber-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb(){
	db,err  := gorm.Open(sqlite.Open("api.db"),&gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database\n",err.Error())
		os.Exit(2)
	}
	log.Println("Database connected")
	db.Logger= logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	Database = DbInstance{Db: db}

}