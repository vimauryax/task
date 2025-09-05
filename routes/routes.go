package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"mv/mvto-do/controller"
	ginSwagger "github.com/swaggo/gin-swagger"
	_"mv/mvto-do/docs"
)

func Routes(c *gin.Engine){
	task := c.Group("/task")

	c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	task.GET("/ping", controller.Ping)
	task.POST("/new", controller.CreateTask)
	task.GET("/:id", controller.GetTaskByIdCont)
	task.GET("/all", controller.GetAllTasksCont)
	task.DELETE("/:id", controller.DeleteTaskByIdCont)
	task.PATCH("/:id", controller.UpdateTaskByIdCont)
}