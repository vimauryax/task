package main

import (
	"fmt"
	_"net/http"
	"mv/mvto-do/routes"
	"github.com/gin-gonic/gin"
	"mv/mvto-do/config"
	"github.com/joho/godotenv"
	"os"
)


func main(){
	err := godotenv.Load(".env")

	if err!=nil{
		fmt.Println("could not read the environment from .env")
	}

	port := os.Getenv("port")
	environment := os.Getenv("GO_ENV")

	if environment == "local"{
		fmt.Println("Running in the local environment")
	}

	config.InitPostgresDBGorm()

	r := gin.Default()

	routes.Routes(r)
	r.Run("localhost:"+port)
}