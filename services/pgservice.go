package services

import (
	"context"
	"errors"
	_"fmt"
	"mv/mvto-do/config"
	"mv/mvto-do/models"
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

	result := config.DB.WithContext(ctx).Create(&task) 
	
	return result.Error
}

func GetTaskById(id string) (*models.Task, error){
	var task models.Task

	result := config.DB.WithContext(ctx).Where("id = ?", id).First(&task)

	return &task, result.Error
}

func GetAllTasks() (*[]models.Task, error){
	var tasks []models.Task

	result := config.DB.WithContext(ctx).Find(&tasks)

	return &tasks, result.Error
}

func DeleteTaskById(id string) error{
	var task models.Task
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
		return errors.New("no fields to update")
	}

	result := config.DB.WithContext(ctx).
		Model(&models.Task{}).
		Where("id = ?", id).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update fields")
	}

	return nil
}
