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
	apihelpers.CustomResponse(c, 200, response)
}

func CreateTask(c *gin.Context) {
	loggerconfig.Info("CreateTask called")
	var response apihelpers.APIRes

	var payload models.Task

	if err := c.ShouldBindJSON(&payload); err != nil {
		loggerconfig.Info("Invalid JSON payload:", err)
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
		apihelpers.SendInternalServerError()
		return
	}
	
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

	task, err := services.GetTaskById(id)

	if err != nil {
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
	apihelpers.CustomResponse(c, 200, response)
}

func GetAllTasksCont(c *gin.Context) {
	tasks, err := services.GetAllTasks()

	var response apihelpers.APIRes

	if err != nil {
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
	apihelpers.CustomResponse(c, 200, response)
}

func DeleteTaskByIdCont(c *gin.Context) {
	var id = c.Param("id")
	var response apihelpers.APIRes

	if err := services.DeleteTaskById(id); err != nil {
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
	apihelpers.CustomResponse(c, 200, response)
}

func UpdateTaskByIdCont(c *gin.Context) {
	id := c.Param("id")

	var response apihelpers.APIRes
	var payload models.TaskUpdatePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
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
	apihelpers.CustomResponse(c, 200, response)
}
