package controller

import (
	_ "encoding/json"
	"mv/mvto-do/loggerconfig"
	"mv/mvto-do/models"
	"mv/mvto-do/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.IndentedJSON(200, gin.H{"success": "connected!"})
}

func CreateTask(c *gin.Context) {
	loggerconfig.Info("CreateTask called")
	var payload models.Task

	if err := c.ShouldBindJSON(&payload); err != nil {
		loggerconfig.Info("Invalid JSON payload:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "failed to fetch the payload",
			"cause": err.Error(),
		})
		return
	}

	if err := services.SaveTask(payload); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "failed to save the task in the database"})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"success": "task created"})
}

func GetTaskByIdCont(c *gin.Context) {
	var id = c.Param("id")

	task, err := services.GetTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "failed to fetch task by id",
			"cause": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"TASK": task})
}

func GetAllTasksCont(c *gin.Context) {
	tasks, err := services.GetAllTasks()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "failed to fetch all tasks",
			"cause": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"Tasks": tasks})
}

func DeleteTaskByIdCont(c *gin.Context) {
	var id = c.Param("id")

	if err := services.DeleteTaskById(id); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "failed to delete the task id" + id,
			"cause": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": "task deleted"})
}

func UpdateTaskByIdCont(c *gin.Context) {
	id := c.Param("id")

	var payload models.TaskUpdatePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON",
			"cause": err.Error(),
		})
		return
	}

	err := services.UpdateTaskById(payload, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to update task",
			"cause": err.Error(),
		})
		return
	}

	task, err := services.GetTaskById(id)

	c.JSON(http.StatusOK, gin.H{"sucess": task})
}
