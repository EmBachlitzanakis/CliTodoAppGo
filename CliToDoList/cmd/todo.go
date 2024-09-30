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
	// Extract titles from the todos
	titles := getTitles(*todos)

	// Build the prompt for the AI model
	result := buildAIPrompt(titles)

	// Initialize the AI model
	llm, err := ollama.New(ollama.WithModel("llama2"))
	if err != nil {
		log.Fatal(err)
	}

	// Call the AI model with the generated prompt
	ctx := context.Background()
	completion, err := llm.Call(ctx, result, llms.WithTemperature(0.8), llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		return nil
	}))

	if err != nil {
		log.Fatal(err)
	}

	// Parse AI completion result to extract priorities
	priorities, err := extractPrioritiesFromAIResponse(completion, (*todos)[0].Title)
	if err != nil {
		log.Fatalf("Failed to extract priorities: %v", err)
	}

	// Render the table of todos
	renderTodosTable(*todos, priorities)
}

// Extract titles from todos list
func getTitles(todos []Todo) []string {
	var titles []string
	for _, todo := range todos {
		titles = append(titles, todo.Title)
	}
	return titles
}

// Build the AI prompt string
func buildAIPrompt(titles []string) string {
	return "Please evaluate the priorities of the following tasks based on their importance. The tasks are separated by commas: " +
		strings.Join(titles, ", ") +
		". For each task, return its priority on a new line using the format: \"Task : Priority\" where Priority is one of the words \"high\", \"medium\", or \"low\". " +
		"Important: Do not include any greetings, explanations, or additional text. Only return the list of tasks with their priorities in the specified format."
}

// Extract priorities from AI response
func extractPrioritiesFromAIResponse(response string, searchTerm string) ([]string, error) {
	startIndex := strings.Index(response, searchTerm)
	if startIndex == -1 {
		return nil, fmt.Errorf("search term not found")
	}

	// Slice the string starting from the search term
	result := response[startIndex:]
	lines := strings.Split(result, "\n")

	// Extract priorities from AI response
	var priorities []string
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			priorities = append(priorities, strings.TrimSpace(parts[1]))
		}
	}

	if len(priorities) == 0 {
		return nil, fmt.Errorf("no priorities found in AI response")
	}

	return priorities, nil
}

// Render the todos table with priorities
func renderTodosTable(todos []Todo, priorities ...[]string) {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Priority", "Created At", "Completed At")

	// Determine if priorities are provided
	var priorityList []string
	if len(priorities) > 0 {
		priorityList = priorities[0]
	}

	for index, t := range todos {
		completed := "❌"
		completedAt := ""
		if t.Completed {
			completed = "✔"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		priority := "N/A"
		if len(priorityList) > index {
			priority = priorityList[index]
		} else {
			priority = t.Priority
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, priority, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}

func (todos *Todos) print() {
	renderTodosTable(*todos)

}
