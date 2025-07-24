package main

import (
	"errors"
	"time"
)

func AddTask(taskName, taskDescription string) error {
	//load json file
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}
	id := tasks.NextAvailableID
	task := Task{
		TaskID:        id,
		TaskName:      taskName,
		Description:   taskDescription,
		CreatedAt:     time.Now(),
		LastUpdatedAt: time.Now(),
		IsCompleted:   false,
	}
	list := tasks.Tasks
	list = append(list, task)

	updatedTasks := Tasks{
		Tasks:           list,
		NextAvailableID: id + 1,
	}

	//save tasks
	return SaveTasks(updatedTasks)
}

func ListTasks() (Tasks, error) {
	return LoadTasks()
}

func DeleteTask(TaskId int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}
	var found bool
	for i, task := range tasks.Tasks {
		if task.TaskID == TaskId {
			found = true
			tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
			break
		}
	}
	if !found {
		return errors.New("given task id does not exist in file")
	}

	//save updated tasks
	return SaveTasks(tasks)
}

func MarkComplete(TaskId int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}
	var found bool
	for i, task := range tasks.Tasks {
		if task.TaskID == TaskId {
			found = true
			if task.IsCompleted {
				return errors.New("given task id already marked")
			}
			tasks.Tasks[i].IsCompleted = true
			break
		}
	}
	if !found {
		return errors.New("given task id does not exist in file")
	}

	//save updated tasks
	return SaveTasks(tasks)
}
