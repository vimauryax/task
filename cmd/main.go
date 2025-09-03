package main

import (
	_"fmt"
	_"net/http"
	"mv/mvto-do/routes"
	"github.com/gin-gonic/gin"
	"mv/mvto-do/config"
)


func main(){
	config.InitPostgresDBGorm()

	r := gin.Default()

	routes.Routes(r)
	r.Run("localhost:8081")
}