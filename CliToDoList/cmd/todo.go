package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	Priority    string
	CreatedAt   time.Time
	CompletedAt *time.Time // Pointer so it can be null

}

type Todos []Todo

func (todos *Todos) add(title string, priority string) error {
	t := *todos
	if err := t.validatePriority(priority); err != nil {
		return err
	}
	todo := Todo{
		Title:       title,
		Completed:   false,
		Priority:    priority,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
	return nil
}

func (todos *Todos) validateIndex(index int) error {

	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func (todos *Todos) validatePriority(priority string) error {

	priorityLevels := []string{"High", "Medium", "Low"}
	if !contains(priorityLevels, priority) {
		err := errors.New("invalid Priority level. Acceptable choices are: High, Medium, Low")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {

	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil
}

func (todos *Todos) editPriority(index int, priority string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	if err := t.validatePriority(priority); err != nil {
		return err
	}

	t[index].Priority = priority

	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	isCompleted := t[index].Completed

	if !isCompleted {
		completeionTime := time.Now()
		t[index].CompletedAt = &completeionTime
	}
	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Priority", "Created At", "Completed At")
	for index, t := range *todos {
		completed := "❌"
		completedAt := ""
		if t.Completed {
			completed = "✔"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.Priority, t.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
