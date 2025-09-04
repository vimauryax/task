package controller

import (
	_ "encoding/json"
	"mv/mvto-do/loggerconfig"
	"mv/mvto-do/apihelpers"
	"mv/mvto-do/models"
	"mv/mvto-do/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	var response apihelpers.APIRes

	if err := services.Ping(); err!=nil{
		loggerconfig.Info("Ping failed, unable to connect")
		apihelpers.SendInternalServerError()
		return 
	}
	response = apihelpers.APIRes{
		Status : true,
			Message : "success",
			ErrorCode : "nil",
			Data :  map[string]interface{}{ 
				"success": "connected",
			},
	}

	loggerconfig.Info("Ping successful!")
	apihelpers.CustomResponse(c, 200, response)
}

func CreateTask(c *gin.Context) {
	var response apihelpers.APIRes

	var payload models.Task

	loggerconfig.Info("trying to create task - CreateTask (controller)")
	if err := c.ShouldBindJSON(&payload); err != nil {
		loggerconfig.Info("Invalid JSON payload : "+err.Error()+" - CreateTask (controller)")
		response = apihelpers.APIRes{
			Status : false,
			Message : "could not read data",
			ErrorCode : "400",
			Data :  map[string]interface{}{ 
				"error": "bad request",
			},
		}

		apihelpers.CustomResponse(c, http.StatusBadRequest, response)
		return
	}

	if err := services.SaveTask(payload); err != nil {
		loggerconfig.Info("Could not save data in the database - CreateTask (controller)")
		apihelpers.SendInternalServerError()
		return
	}
	
	loggerconfig.Info("Task saved in the database - CreateTask (controller)")
	response = apihelpers.APIRes{
		Status : true,
			Message : "success",
			ErrorCode : "nil",
			Data :  map[string]interface{}{ 
				"success": "task created",
			},
	}
	apihelpers.CustomResponse(c, 200, response)
}

func GetTaskByIdCont(c *gin.Context) {
	var id = c.Param("id")
	var response apihelpers.APIRes

	loggerconfig.Info("Fetching task for id : "+id+" - GetTaskById (controller)")
	task, err := services.GetTaskById(id)

	if err != nil {
		loggerconfig.Info("failed to fetch task from database - GetTaskById (controller)")
		apihelpers.SendInternalServerError()
		return
	}
	response = apihelpers.APIRes{
		Status : true,
			Message : "success",
			ErrorCode : "nil",
			Data :  map[string]interface{}{ 
				"task": task,
			},
	}
	loggerconfig.Info("fetched task from database - GetTaskById (controller)")
	apihelpers.CustomResponse(c, 200, response)
}

func GetAllTasksCont(c *gin.Context) {

	loggerconfig.Info("fetching all tasks from database - GetAllTasks (controller)")
	tasks, err := services.GetAllTasks()

	var response apihelpers.APIRes

	if err != nil {
		loggerconfig.Info("failed to fetch tasks from database - GetAllTasks (controller)")
		apihelpers.SendInternalServerError()
		return
	}

	
	response = apihelpers.APIRes{
		Status : true,
			Message : "success",
			ErrorCode : "nil",
			Data :  map[string]interface{}{ 
				"tasks": tasks,
			},
	}
	loggerconfig.Info("fetched all tasks from database - GetAllTasks (controller)")
	apihelpers.CustomResponse(c, 200, response)
}

func DeleteTaskByIdCont(c *gin.Context) {
	var id = c.Param("id")
	var response apihelpers.APIRes
	loggerconfig.Info("Deleting task from database with id : "+id+" - DeleteTaskById (controller)")
	if err := services.DeleteTaskById(id); err != nil {
		loggerconfig.Info("failed to delete task from database with id : "+id+" - DeleteTaskById (controller)")
		apihelpers.SendInternalServerError()
		return
	}

	response = apihelpers.APIRes{
		Status : true,
			Message : "success",
			ErrorCode : "nil",
			Data :  map[string]interface{}{ 
				"success": "task deleted",
			},
	}

	loggerconfig.Info("Deleted task from database with id : "+id+" - DeleteTaskById (controller)")
	apihelpers.CustomResponse(c, 200, response)
}

func UpdateTaskByIdCont(c *gin.Context) {
	id := c.Param("id")

	loggerconfig.Info("Updating task in database with id : "+id+" - UpdateTaskById (controller)")
	var response apihelpers.APIRes
	var payload models.TaskUpdatePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		loggerconfig.Info("failed bind with json : "+id+" -UpdateTaskById (controller)")
		response = apihelpers.APIRes{
			Status : false,
			Message : "invalid json",
			ErrorCode : "400",
			Data :  map[string]interface{}{ 
				"error": "bad request",
			},
		}
		
		apihelpers.CustomResponse(c, http.StatusBadRequest, response)
		return
	}

	err := services.UpdateTaskById(payload, id)
	if err != nil {
		loggerconfig.Info("failed to update task in database with id : "+id+" - UpdateTaskById (controller)")
		apihelpers.SendInternalServerError()
		return
	}

	task, _ := services.GetTaskById(id)

	response = apihelpers.APIRes{
		Status : true,
			Message : "success",
			ErrorCode : "nil",
			Data :  map[string]interface{}{ 
				"task": task,
			},
	}

	loggerconfig.Info("updated task in database with id : "+id+" - UpdateTaskById (controller)")
	apihelpers.CustomResponse(c, 200, response)
}
