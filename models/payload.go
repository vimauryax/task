package models

import (
	"time"
)

type Task struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Title          string    `json:"title" binding:"required"`
	Desc           string    `json:"desc"`
	TimestampBegin time.Time `json:"timestampBegin" binding:"required"`
	TimestampEnd   time.Time `json:"timestampEnd" binding:"required"`
	Status         string    `json:"status" binding:"required,oneof=pending done outdated"`
}


type TaskUpdatePayload struct {
	Title        *string    `json:"title,omitempty"`
	Desc         *string    `json:"desc,omitempty"`
	Status       *string    `json:"status,omitempty"`
	TimestampEnd *time.Time `json:"timestampEnd,omitempty"`
}