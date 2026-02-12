package main

import (
	"fmt"
	"time"
)

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

type Task interface {
	GetID() int
	GetTitle() string
	GetDescription() string
	GetPriority() Priority
	GetCreatedAt() time.Time
	IsCompleted() bool
	Complete()
	Display() string
}

type BaseTask struct {
	ID          int
	Title       string
	Description string
	Priority    Priority
	CreatedAt   time.Time
	Completed   bool
}

func (b *BaseTask) GetID() int {
	return b.ID
}

func (b *BaseTask) GetTitle() string {
	return b.Title
}

func (b *BaseTask) GetDescription() string {
	return b.Description
}

func (b *BaseTask) GetPriority() Priority {
	return b.Priority
}

func (b *BaseTask) GetCreatedAt() time.Time {
	return b.CreatedAt
}

func (b *BaseTask) IsCompleted() bool {
	return b.Completed
}

func (b *BaseTask) Complete() {
	b.Completed = true
}

type WorkTask struct {
	BaseTask
	Project  string
	Deadline time.Time
}

func (w WorkTask) Display() string {
	status := "Pending"
	if w.Completed {
		status = "Completed"
	}
	return fmt.Sprintf("[WORK] ID: %d | %s | Priority: %s | Status: %s\nProject: %s | Deadline: %s\n%s",
		w.ID, w.Title, w.Priority, status, w.Project, w.Deadline.Format("2006-01-02"), w.Description)
}

type PersonalTask struct {
	BaseTask
	Category string
}

func (p PersonalTask) Display() string {
	status := "Pending"
	if p.Completed {
		status = "Completed"
	}
	return fmt.Sprintf("[PERSONAL] ID: %d | %s | Priority: %s | Status: %s\nCategory: %s\n%s",
		p.ID, p.Title, p.Priority, status, p.Category, p.Description)
}

type ShoppingTask struct {
	BaseTask
	Items  []string
	Budget float64
}

func (s ShoppingTask) Display() string {
	status := "Pending"
	if s.Completed {
		status = "Completed"
	}
	itemList := ""
	for i, item := range s.Items {
		itemList += fmt.Sprintf("\n  %d. %s", i+1, item)
	}
	return fmt.Sprintf("[SHOPPING] ID: %d | %s | Priority: %s | Status: %s\nBudget: $%.2f\nItems:%s",
		s.ID, s.Title, s.Priority, status, s.Budget, itemList)
}

type TaskManager struct {
	tasks  []Task
	nextID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:  make([]Task, 0),
		nextID: 1,
	}
}

func (tm *TaskManager) AddTask(task Task) {
	tm.tasks = append(tm.tasks, task)
	tm.nextID++
}

func (tm *TaskManager) GetTask(id int) (Task, error) {
	for _, task := range tm.tasks {
		if task.GetID() == id {
			return task, nil
		}
	}
	return nil, fmt.Errorf("task with ID %d not found", id)
}

func (tm *TaskManager) GetAllTasks() []Task {
	return tm.tasks
}

func (tm *TaskManager) GetTasksByPriority(p Priority) []Task {
	filtered := make([]Task, 0)
	for _, task := range tm.tasks {
		if task.GetPriority() == p {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

func (tm *TaskManager) GetCompletedTasks() []Task {
	completed := make([]Task, 0)
	for _, task := range tm.tasks {
		if task.IsCompleted() {
			completed = append(completed, task)
		}
	}
	return completed
}

func (tm *TaskManager) GetPendingTasks() []Task {
	pending := make([]Task, 0)
	for _, task := range tm.tasks {
		if !task.IsCompleted() {
			pending = append(pending, task)
		}
	}
	return pending
}

func (tm *TaskManager) CompleteTask(id int) error {
	task, err := tm.GetTask(id)
	if err != nil {
		return err
	}
	task.Complete()
	return nil
}

func (tm *TaskManager) DisplayAllTasks() {
	if len(tm.tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	for _, task := range tm.tasks {
		fmt.Println(task.Display())
		fmt.Println("---")
	}
}

func main() {
	manager := NewTaskManager()

	workTask := &WorkTask{
		BaseTask: BaseTask{
			ID:          manager.nextID,
			Title:       "Complete API Documentation",
			Description: "Write comprehensive API docs for the new endpoints",
			Priority:    High,
			CreatedAt:   time.Now(),
			Completed:   false,
		},
		Project:  "Backend Redesign",
		Deadline: time.Now().AddDate(0, 0, 7),
	}
	manager.AddTask(workTask)

	personalTask := &PersonalTask{
		BaseTask: BaseTask{
			ID:          manager.nextID,
			Title:       "Morning Workout",
			Description: "30 minutes cardio + strength training",
			Priority:    Medium,
			CreatedAt:   time.Now(),
			Completed:   false,
		},
		Category: "Health",
	}
	manager.AddTask(personalTask)

	shoppingTask := &ShoppingTask{
		BaseTask: BaseTask{
			ID:          manager.nextID,
			Title:       "Weekly Groceries",
			Description: "Buy groceries for the week",
			Priority:    Low,
			CreatedAt:   time.Now(),
			Completed:   false,
		},
		Items:  []string{"Milk", "Eggs", "Bread", "Vegetables", "Fruits"},
		Budget: 150.00,
	}
	manager.AddTask(shoppingTask)

	urgentTask := &WorkTask{
		BaseTask: BaseTask{
			ID:          manager.nextID,
			Title:       "Fix Production Bug",
			Description: "Critical bug in payment processing",
			Priority:    Urgent,
			CreatedAt:   time.Now(),
			Completed:   false,
		},
		Project:  "Payment System",
		Deadline: time.Now().AddDate(0, 0, 1),
	}
	manager.AddTask(urgentTask)

	fmt.Println("=== ALL TASKS ===")
	manager.DisplayAllTasks()

	fmt.Println("\n=== COMPLETING TASK 2 ===")
	manager.CompleteTask(2)

	fmt.Println("\n=== URGENT PRIORITY TASKS ===")
	urgentTasks := manager.GetTasksByPriority(Urgent)
	for _, task := range urgentTasks {
		fmt.Println(task.Display())
		fmt.Println("---")
	}

	fmt.Println("\n=== COMPLETED TASKS ===")
	completedTasks := manager.GetCompletedTasks()
	fmt.Printf("Total completed: %d\n", len(completedTasks))
	for _, task := range completedTasks {
		fmt.Println(task.Display())
		fmt.Println("---")
	}

	fmt.Println("\n=== PENDING TASKS ===")
	pendingTasks := manager.GetPendingTasks()
	fmt.Printf("Total pending: %d\n", len(pendingTasks))
	for _, task := range pendingTasks {
		fmt.Println(task.Display())
		fmt.Println("---")
	}
}
