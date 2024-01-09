package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	utils "github.com/Mix-Liten/go-todo_cli_app/utils"
	"github.com/alexeyco/simpletable"
)

// item represents a single todo item with task details and completion status
type item struct {
	Task        string    // Task description
	Done        bool      // Completion status (true if done)
	CreatedAt   time.Time // Time when the task was created
	CompletedAt time.Time // Time when the task was completed
}

// Todos is a slice of todo items
type Todos []item

// Add adds a new task to the list of todos
func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

// Complete marks a task as completed by its index in the list of todos
func (t *Todos) Complete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

// Delete removes a task by its index from the list of todos
func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

// Load reads the todos from a file and decodes the JSON data
func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

// Store encodes the todos to JSON and writes them to a file
func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// Print generates an ASCII table of the todo list and prints it to the console
func (t *Todos) Print() {
	// Creating an ASCII table
	table := simpletable.New()

	// Setting the table header
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell
	for idx, item := range *t {
		task := utils.Blue(item.Task)
		done := utils.Blue("no")
		if item.Done {
			task = utils.Green(fmt.Sprintf("\u2705 %s", item.Task))
			done = utils.Green("yes")
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx+1)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})
	}
	table.Body = &simpletable.Body{Cells: cells}

	// Setting the footer of the table
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: utils.Red(fmt.Sprintf("You have %d pending todos", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}

// CountPending counts the number of pending tasks in the todo list
func (t *Todos) CountPending() int {
	total := 0
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}

	return total
}
