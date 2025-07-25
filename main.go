package main

import (
	"flag"
	"fmt"
)

func main() {
	//first define cl flags
	add := flag.String("add", "", "The task you want to add in your TODO list")
	desc := flag.String("desc", "", "Description of the task")
	list := flag.Bool("list", false, "List of all TODO tasks")
	delete := flag.Int("delete", 0, "Delete the given task id")
	complete := flag.Int("complete", 0, "Marked true for the given task id")

	//parse flags
	flag.Parse()

	if *add != "" {
		//add a task
		err := AddTask(*add, *desc)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if *list {
		tasks, err := ListTasks()
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(tasks.Tasks) == 0 {
			fmt.Println("No tasks found")
			return
		}
		for _, task := range tasks.Tasks {
			status := "Pending"
			if task.IsCompleted {
				status = "Completed"
			}
			fmt.Printf(" %d | %s  | Status: %s\n", task.TaskID, task.TaskName, status)
		}
	}

	if *delete > 0 {
		err := DeleteTask(*delete)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Your task id:%d has been deleted\n", *delete)
	}
	if *complete > 0 {
		err := MarkComplete(*complete)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Your task id:%d has been marked as completed\n", *complete)
	}

	if *add == "" && *list == false && *delete < 1 && *complete < 1 {
		fmt.Println("\nPlease use -help for available commands\n")
	}

}
