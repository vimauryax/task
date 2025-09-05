package main

import (
	"fmt"
	_"net/http"
	"mv/mvto-do/routes"
	"github.com/gin-gonic/gin"
	"mv/mvto-do/config"
	"github.com/joho/godotenv"
	"os"
	"mv/mvto-do/loggerconfig"
)


func main(){
	loggerconfig.LogrusInitialize()
	err := godotenv.Load(".env")

	if err!=nil{
		fmt.Println("could not read the environment from .env")
	}

	port := os.Getenv("port")
	environment := os.Getenv("GO_ENV")

	if environment == "local"{
		fmt.Println("Running in local environment")
	} else if environment == "prod"{
		fmt.Println("Running in prod environment")
	} else {
		fmt.Println("Select the right environment")
	}

	config.InitPostgresDBGorm(environment)

	r := gin.Default()

	routes.Routes(r)
	r.Run("localhost:"+port)
}
