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
		loggerconfig.Info("Ping (controller) - Ping failed, unable to connect")
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

	loggerconfig.Info("Ping (controller) - Ping successful!")
	apihelpers.CustomResponse(c, http.StatusOK, response)
}

func CreateTask(c *gin.Context) {
	var response apihelpers.APIRes

	var payload models.Task

	loggerconfig.Info("CreateTask (controller) - trying to create task")
	if err := c.ShouldBindJSON(&payload); err != nil {
		loggerconfig.Info("CreateTask (controller) - Invalid JSON payload : "+err.Error())
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
		loggerconfig.Info("CreateTask (controller) - Could not save data in the database")
		apihelpers.SendInternalServerError()
		return
	}
	
	loggerconfig.Info("CreateTask (controller) - Task saved in the database")
	response = apihelpers.APIRes{
		Status : true,
			Message : "success",
			ErrorCode : "nil",
			Data :  map[string]interface{}{ 
				"success": "task created",
			},
	}
	apihelpers.CustomResponse(c, http.StatusOK, response)
}

func GetTaskByIdCont(c *gin.Context) {
	var id = c.Param("id")
	var response apihelpers.APIRes

	loggerconfig.Info("GetTaskById (controller) - Fetching task for id : "+id)
	task, err := services.GetTaskById(id)

	if err != nil {
		loggerconfig.Info("GetTaskById (controller) - failed to fetch task from database")
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
	loggerconfig.Info("GetTaskById (controller) - fetched task from database")
	apihelpers.CustomResponse(c, http.StatusOK, response)
}

// @Tags space cmots V2
// @Description Create Task
// @Success 200 {object} apihelpers.APIRes{data=[]models.Task}
// @Failure 400 {object} apihelpers.APIRes
// @Router /api/task/all [GET]
func GetAllTasksCont(c *gin.Context) {

	loggerconfig.Info("GetAllTasks (controller) - fetching all tasks from database")
	tasks, err := services.GetAllTasks()

	var response apihelpers.APIRes

	if err != nil {
		loggerconfig.Info("GetAllTasks (controller) - failed to fetch tasks from database")
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
	loggerconfig.Info("GetAllTasks (controller) - fetched all tasks from database")
	apihelpers.CustomResponse(c, http.StatusOK, response)
}

func DeleteTaskByIdCont(c *gin.Context) {
	var id = c.Param("id")
	var response apihelpers.APIRes
	loggerconfig.Info("DeleteTaskById (controller) - Deleting task from database with id : "+id)
	if err := services.DeleteTaskById(id); err != nil {
		loggerconfig.Info("DeleteTaskById (controller) - failed to delete task from database with id : "+id)
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

	loggerconfig.Info("DeleteTaskById (controller) - Deleted task from database with id : "+id)
	apihelpers.CustomResponse(c, http.StatusOK, response)
}

func UpdateTaskByIdCont(c *gin.Context) {
	id := c.Param("id")

	loggerconfig.Info("UpdateTaskById (controller) - Updating task in database with id : "+id)
	var response apihelpers.APIRes
	var payload models.TaskUpdatePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		loggerconfig.Info("UpdateTaskById (controller) - failed bind with json : "+id)
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
		loggerconfig.Info("UpdateTaskById (controller) - failed to update task in database with id : "+id)
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

	loggerconfig.Info("UpdateTaskById (controller) - updated task in database with id : "+id)
	apihelpers.CustomResponse(c, http.StatusOK, response)
}
