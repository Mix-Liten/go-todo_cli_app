package todo

import (
	"os"
	"reflect"
	"testing"

	todo "github.com/Mix-Liten/go-todo_cli_app/core/todo"
)

func TestAdd(t *testing.T) {
	// Create a new Todos slice
	todos := &todo.Todos{}

	// Add a task
	task := "Test Task"
	todos.Add(task)

	// Verify the task is added correctly
	if len(*todos) != 1 || (*todos)[0].Task != task { // Corrected slice access
		t.Errorf("Add() failed, expected: %s, got: %v", task, todos)
	}
}

func TestComplete(t *testing.T) {
	// Create a new Todos slice
	todos := &todo.Todos{}
	task := "Test Task"
	todos.Add(task)

	// Complete the added task
	err := todos.Complete(1)

	// Verify the task is marked as completed without errors
	if err != nil || !(*todos)[0].Done { // Corrected slice access
		t.Errorf("Complete() failed, expected: true, got: %v, error: %v", (*todos)[0].Done, err)
	}
}

func TestDelete(t *testing.T) {
	// Create a new Todos slice
	todos := &todo.Todos{}
	tasks := []string{"Task 1", "Task 2", "Task 3"}
	for _, task := range tasks {
		todos.Add(task)
	}

	// Delete a task
	err := todos.Delete(2)

	// Verify the task is deleted without errors
	expectedTasks := []string{"Task 1", "Task 3"}
	if err != nil || !reflect.DeepEqual(getTaskNames(*todos), expectedTasks) { // Corrected slice access
		t.Errorf("Delete() failed, expected: %v, got: %v, error: %v", expectedTasks, getTaskNames(*todos), err)
	}
}

func TestLoadAndStore(t *testing.T) {
	// Create test data
	testFilename := "test-todos.json"
	defer os.Remove(testFilename)

	// Create a new Todos slice and add tasks
	todos := &todo.Todos{}
	tasks := []string{"Task 1", "Task 2", "Task 3"}
	for _, task := range tasks {
		todos.Add(task)
	}

	// Store tasks to a file
	err := todos.Store(testFilename)
	if err != nil {
		t.Fatalf("Store() failed: %v", err)
	}

	// Load tasks from the file
	loadedTodos := &todo.Todos{}
	err = loadedTodos.Load(testFilename)
	if err != nil {
		t.Fatalf("Load() failed: %v", err)
	}

	// Verify loaded tasks match the original tasks
	if !reflect.DeepEqual(getTaskNames(*loadedTodos), tasks) { // Corrected slice access
		t.Errorf("Load/Store operations failed, expected: %v, got: %v", tasks, getTaskNames(*loadedTodos))
	}
}

func getTaskNames(t todo.Todos) []string { // Corrected type reference
	names := make([]string, len(t))
	for i, task := range t {
		names[i] = task.Task
	}
	return names
}
