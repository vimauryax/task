package services

import (
	"context"
	"errors"
	_"fmt"
	"mv/mvto-do/config"
	"mv/mvto-do/models"
	"mv/mvto-do/loggerconfig"
)
var ctx context.Context

func Ping() error{
	return nil
}

func SaveTask(payload models.Task) error {
	task := models.Task{
		Title:          payload.Title,
		Desc:           payload.Desc,
		TimestampBegin: payload.TimestampBegin,
		TimestampEnd:   payload.TimestampEnd,
		Status:         payload.Status,
	}

	loggerconfig.Info("SaveTask (services) - saving data in the database")
	result := config.DB.WithContext(ctx).Create(&task) 
	
	return result.Error
}

func GetTaskById(id string) (*models.Task, error){
	var task models.Task

	loggerconfig.Info("GetTaskById (services) - fetching task from database")
	result := config.DB.WithContext(ctx).Where("id = ?", id).First(&task)

	return &task, result.Error
}

func GetAllTasks() (*[]models.Task, error){
	var tasks []models.Task

	loggerconfig.Info("GetAllTasks (services) - fetching all tasks from database")
	result := config.DB.WithContext(ctx).Find(&tasks)

	return &tasks, result.Error
}

func DeleteTaskById(id string) error{
	var task models.Task

	loggerconfig.Info("DeleteTaskById (services) - deleting task from database with id : "+id)
	result := config.DB.WithContext(ctx).Where("id = ?", id).Delete(&task)
	return result.Error
}

func UpdateTaskById(payload models.TaskUpdatePayload, id string) error {
	updates := make(map[string]interface{})

	if payload.Title != nil {
		updates["title"] = *payload.Title
	}
	if payload.Desc != nil {
		updates["desc"] = *payload.Desc
	}
	if payload.Status != nil {
		updates["status"] = *payload.Status
	}
	if payload.TimestampEnd != nil {
		updates["timestamp_end"] = *payload.TimestampEnd
	}

	if len(updates) == 0 {
		loggerconfig.Info("UpdateTaskById (services) - no fields to update")
		return errors.New("no fields to update")
	}

	result := config.DB.WithContext(ctx).
		Model(&models.Task{}).
		Where("id = ?", id).
		Updates(updates)

	if result.Error != nil {
		loggerconfig.Info("UpdateTaskById (services) - failed to update in the database")
		return result.Error
	}
	if result.RowsAffected == 0 {
		loggerconfig.Info("UpdateTaskById (services) - 0 rows affected.")
		return errors.New("failed to update fields")
	}

	return nil
}
