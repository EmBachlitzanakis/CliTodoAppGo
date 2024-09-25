package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aquasecurity/table"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
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

func (todos *Todos) Aiprint() {
	var titles []string
	for _, todo := range *todos {
		titles = append(titles, todo.Title)
	}
	//result := " evaluate the priority of those tasks i gave you separated by ','  The list is :" + strings.Join(titles, ", ") + " . Between three options : high, medium, sort. Return a list with their priorities. example of output i expect if i have 4 items : high,high,low,medium. IMPORTANT i want to be return just the priorities ,no greetings or anything else"

	result := "Please evaluate the priorities of the following tasks based on their importance. The tasks are separated by commas: " +
		strings.Join(titles, ", ") +
		". For each task, return its priority on a new line using the format: \"Task : Priority\" where Priority is one of the words \"high\", \"medium\", or \"low\". " +
		"Important: Do not include any greetings, explanations, or additional text. Only return the list of tasks with their priorities in the specified format."
	llm, err := ollama.New(ollama.WithModel("llama2"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	completion, err := llm.Call(ctx, result,
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {

			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(completion)
	searchTerm := (*todos)[0].Title

	// Find the index of "Buy milk"
	startIndex := strings.Index(completion, searchTerm)
	values := make([]string, 0)
	// // Check if the search term is found
	if startIndex != -1 {
		// Slice the string starting from "Buy milk"
		result := completion[startIndex:]
		// Split the result into lines
		lines := strings.Split(result, "\n")

		// Create a slice to hold the values after the colon

		// Iterate through each line
		for _, line := range lines {
			// Split each line by the colon
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				// Trim spaces and append the value to the slice
				values = append(values, strings.TrimSpace(parts[1]))
			}
		}

		// Print the resulting slice
		//fmt.Println(values)
		//fmt.Println(len(values))
	} else {
		fmt.Println("Search term not found!")
	}

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

		if index < len(values) {
			table.AddRow(strconv.Itoa(index), t.Title, completed, values[index], t.CreatedAt.Format(time.RFC1123), completedAt)
		} else {
			// Handle case where there are not enough values
			table.AddRow(strconv.Itoa(index), t.Title, completed, "N/A", t.CreatedAt.Format(time.RFC1123), completedAt)
		}
		//table.AddRow(strconv.Itoa(index), t.Title, completed, values[1], t.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
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
