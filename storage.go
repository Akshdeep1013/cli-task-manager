package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadTasks() (Tasks, error) {
	filePath := "/home/deep/personal-projects/cli-task-manager/tasks.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return Tasks{Tasks: []Task{}, NextAvailableID: 1}, nil
		}
		return Tasks{}, err
	}
	var tasks Tasks
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error while parsing data: ", err)
		return Tasks{}, err
	}
	return tasks, nil
}

func SaveTasks(updatedTasks Tasks) error {
	jsonBytes, err := json.Marshal(updatedTasks)
	if err != nil {
		fmt.Println("Error while parsing data: ", err)
		return err
	}

	filePath := "/home/deep/personal-projects/cli-task-manager/tasks.json"
	err = os.WriteFile(filePath, jsonBytes, 0644)
	if err != nil {
		fmt.Println("Error while writting into the file: ", err)
		return err
	}
	return nil
}
