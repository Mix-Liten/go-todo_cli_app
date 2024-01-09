package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/Mix-Liten/go-todo_cli_app/core/todo"
)

const (
	todoFile = ".todos.json" // Defining the file name for storing todo items
)

func main() {
	add := flag.Bool("add", false, "add a new todo")                // Flag to add a new todo
	complete := flag.Int("complete", 0, "mark a todo as completed") // Flag to mark a todo as completed
	delete := flag.Int("delete", 0, "delete a todo")                // Flag to delete a todo
	list := flag.Bool("list", false, "list all todos")              // Flag to list all todos
	flag.Parse()                                                    // Parsing the command-line flags

	todos := &todo.Todos{} // Initializing a Todos struct from the custom package

	// Loading existing todos from the file
	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error()) // Printing error message to standard error output
		os.Exit(1)                           // Exiting the program with an error status
	}

	// Performing different operations based on the provided command-line flags
	switch {
	case *add: // Adding a new todo
		task, err := getInput(os.Stdin, flag.Args()...) // Getting input for a new task
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error()) // Printing error message to standard error output
			os.Exit(1)                           // Exiting the program with an error status
		}

		todos.Add(task)             // Adding the task to the todos list
		err = todos.Store(todoFile) // Storing the updated todos to a file
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error()) // Printing error message to standard error output
			os.Exit(1)                           // Exiting the program with an error status
		}
	case *complete > 0: // Marking a todo as completed
		err := todos.Complete(*complete) // Completing a todo based on the provided index
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error()) // Printing error message to standard error output
			os.Exit(1)                           // Exiting the program with an error status
		}
		err = todos.Store(todoFile) // Storing the updated todos to a file
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error()) // Printing error message to standard error output
			os.Exit(1)                           // Exiting the program with an error status
		}
	case *delete > 0: // Deleting a todo
		err := todos.Delete(*delete) // Deleting a todo based on the provided index
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error()) // Printing error message to standard error output
			os.Exit(1)                           // Exiting the program with an error status
		}
		err = todos.Store(todoFile) // Storing the updated todos to a file
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error()) // Printing error message to standard error output
			os.Exit(1)                           // Exiting the program with an error status
		}
	case *list: // Listing all todos
		todos.Print() // Printing all the todos
	default:
		fmt.Fprintln(os.Stdout, "invalid command") // Printing a message to standard output for an invalid command
		os.Exit(0)                                 // Exiting the program with a success status
	}
}

// getInput retrieves the input text from the user or command-line arguments
func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil // Returning joined input arguments
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", nil // Returning an empty string if there's an error
	}

	text := scanner.Text()
	if len(text) == 0 {
		return "", errors.New("empty todo is not allowed") // Returning an error for an empty todo
	}

	return text, nil // Returning the input text and no error
}
