package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	"mv/mvto-do/models"
	"log"
)

var DB *gorm.DB

func InitPostgresDBGorm(){
	dsn := "host=localhost user=postgres password=tiger dbname=to-do port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!=nil{
		fmt.Println("failed to connected to the database")
	}
	DB = database

	err = DB.AutoMigrate(&models.Task{})
    if err != nil {
        log.Fatal("Auto migration failed:", err)
    }
}