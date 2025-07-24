package main

import "time"

type Task struct {
	TaskID        int       `json:"task_id"` //make primary key
	CreatedAt     time.Time `json:"created_at"`
	TaskName      string    `json:"task_name"`
	Description   string    `json:"description"`
	IsCompleted   bool      `json:"is_completed"`
	LastUpdatedAt time.Time `json:"last_updated_at"`
}

type Tasks struct {
	Tasks           []Task `json:"tasks"`
	NextAvailableID int    `json:"next_available_id"`
}
