package main

import (
	"fmt"
	"time"
)

// Priority levels for tasks
type Priority int

const (
	Low Priority = iota
	Medium
	High
	Urgent
)

func (p Priority) String() string {
	switch p {
	case Low:
		return "Low"
	case Medium:
		return "Medium"
	case High:
		return "High"
	case Urgent:
		return "Urgent"
	default:
		return "Unknown"
	}
}

// TODO: Create a Task interface with the following methods:
// - GetID() int
// - GetTitle() string
// - GetDescription() string
// - GetPriority() Priority
// - GetCreatedAt() time.Time
// - IsCompleted() bool
// - Complete()
// - Display() string

// TODO: Create a BaseTask struct with common fields:
// - ID (int)
// - Title (string)
// - Description (string)
// - Priority (Priority)
// - CreatedAt (time.Time)
// - Completed (bool)

// TODO: Implement getter methods for BaseTask

// TODO: Create WorkTask struct that embeds BaseTask and adds:
// - Project (string)
// - Deadline (time.Time)

// TODO: Create PersonalTask struct that embeds BaseTask and adds:
// - Category (string) - e.g., "Health", "Learning", "Hobby"

// TODO: Create ShoppingTask struct that embeds BaseTask and adds:
// - Items ([]string)
// - Budget (float64)

// TODO: Implement Display() method for each task type with custom formatting

// TODO: Create a TaskManager struct with:
// - tasks ([]Task)
// - nextID (int)

// TODO: Implement the following methods for TaskManager:
// - AddTask(task Task) - adds a task to the list
// - GetTask(id int) (Task, error) - retrieves a task by ID
// - GetAllTasks() []Task - returns all tasks
// - GetTasksByPriority(p Priority) []Task - filters by priority
// - GetCompletedTasks() []Task - returns completed tasks
// - GetPendingTasks() []Task - returns pending tasks
// - CompleteTask(id int) error - marks a task as completed
// - DisplayAllTasks() - prints all tasks

func main() {
	// TODO: Create a TaskManager instance
	// TODO: Add various types of tasks (Work, Personal, Shopping)
	// TODO: Display all tasks
	// TODO: Complete some tasks
	// TODO: Display tasks by priority
	// TODO: Display completed vs pending tasks

	fmt.Println("Build your Task Manager here!")
	fmt.Println("\nExpected features:")
	fmt.Println("1. Create different types of tasks")
	fmt.Println("2. Add tasks to manager")
	fmt.Println("3. Complete tasks")
	fmt.Println("4. Filter by priority")
	fmt.Println("5. View completed/pending tasks")
}
