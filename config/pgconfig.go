package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	"mv/mvto-do/models"
	"log"
)

var DB *gorm.DB

func InitPostgresDBGorm(env string){
	dsn := "host=localhost user=postgres password=tiger dbname=to-do port=5432 sslmode=disable"
	dsn1 := "host=localhost user=postgres password=tiger dbname=to-do-prod port=5432 sslmode=disable"

	var fdsn string

	if env == "local"{
		fdsn = dsn
	} else if env == "prod"{
		fdsn = dsn1
	}
	database, err := gorm.Open(postgres.Open(fdsn), &gorm.Config{})
	if err!=nil{
		fmt.Println("failed to connected to the database")
	}
	DB = database

	err = DB.AutoMigrate(&models.Task{})
    if err != nil {
        log.Fatal("Auto migration failed:", err)
    }
}