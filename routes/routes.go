package routes

import (
	"github.com/gin-gonic/gin"
	"mv/mvto-do/controller"
)

func Routes(c *gin.Engine){
	task := c.Group("/task")

	task.GET("/ping", controller.Ping)
	task.POST("/new", controller.CreateTask)
	task.GET("/:id", controller.GetTaskByIdCont)
	task.GET("/all", controller.GetAllTasksCont)
	task.DELETE("/:id", controller.DeleteTaskByIdCont)
	task.PATCH("/:id", controller.UpdateTaskByIdCont)
}