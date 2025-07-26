package main

import (
	"testing"
)

func TestAddTask(t *testing.T) {
	fakeStorage := MockStorage{
		FakeData: Tasks{
			Tasks:           []Task{},
			NextAvailableID: 1,
		},
	}
	err := AddTask("TEST1", "Watch test match often in Cinema", &fakeStorage)
	if err != nil {
		t.Error("AddTask should not return an error, but got:", err)
	}
	//verify whether it is added or not
	tasks, err := fakeStorage.Load()
	if err != nil {
		t.Error("Could not load task from mock:", err)
		return
	}

	if len(tasks.Tasks) != 1 {
		t.Error("Expected 1 task but got:", len(tasks.Tasks))
		return
	}

	task := tasks.Tasks[0]
	if task.TaskName != "TEST1" {
		t.Error("Expected Task Name: 'TEST1' but got:", task.TaskName)
		return
	}

	if task.TaskID != 1 {
		t.Error("Expected TaskID: 1 but got:", task.TaskID)
		return
	}

	t.Log("AddTask test completed successfully - no real files touched!")
}

func TestListTasks(t *testing.T) {
	fakeStorage := MockStorage{
		FakeData: Tasks{
			Tasks:           []Task{},
			NextAvailableID: 1,
		},
	}
	tasks, err := ListTasks(&fakeStorage)

	if err != nil {
		t.Error("ListTasks should not return an error, but got:", err)
	}

	if len(tasks.Tasks) != 0 {
		t.Error("Expected 0 task but got:", len(tasks.Tasks))
		return
	}

	t.Log("ListTasks test completed successfully - no real files touched!")

}

func TestDeleteTask(t *testing.T) {
	fakeStorage := MockStorage{
		FakeData: Tasks{
			Tasks:           []Task{},
			NextAvailableID: 1,
		},
	}
	//lets add dummy test task
	err := AddTask("TEST1", "Watch test match often in Cinema", &fakeStorage)
	if err != nil {
		t.Error("AddTask should not return an error, but got:", err)
	}

	err = DeleteTask(1, &fakeStorage)
	if err != nil {
		t.Error("DeleteTask should not return an error")
	}

	//now verify
	tasks, err := fakeStorage.Load()
	if err != nil {
		t.Error("Could not load task from mock:", err)
		return
	}

	if len(tasks.Tasks) != 0 {
		t.Error("Expected 0 task but got:", len(tasks.Tasks))
		return
	}

	t.Log("DeleteTask test completed successfully")
}

func TestMarkComplete(t *testing.T) {
	fakeStorage := MockStorage{
		FakeData: Tasks{
			Tasks:           []Task{},
			NextAvailableID: 1,
		},
	}
	//lets add dummy test task
	err := AddTask("TEST1", "Watch test match often in Cinema", &fakeStorage)
	if err != nil {
		t.Error("AddTask should not return an error, but got:", err)
	}

	err = MarkComplete(1, &fakeStorage)
	if err != nil {
		t.Error("MarkComplete should not return an error")
	}

	//verfiy whether it is completed or not
	tasks, err := fakeStorage.Load()
	if err != nil {
		t.Error("Could not load task from mock:", err)
		return
	}

	if len(tasks.Tasks) != 1 {
		t.Error("Expected 1 task but got:", len(tasks.Tasks))
		return
	}

	task := tasks.Tasks[0]

	if task.TaskID != 1 {
		t.Error("Expected TaskID: 1 but got:", task.TaskID)
		return
	}

	if !task.IsCompleted {
		t.Error("Expected Task is completed but not completd at its status :", task.IsCompleted)
		return
	}

	t.Log("MarkComplete test completed successfully")
}
