# go-todo_cli_app

This is a simple command-line application for managing a todo list. It provides functionalities to add, complete, delete, and list tasks. The todo items are stored in a JSON file.

## Installation

To use this application, you need Go installed on your system.

```bash
go get -u github.com/Mix-Liten/go-todo_cli_app
```

## usage
To interact with the todo application, you can use the command line with the following options:

- List all todos:
  ```bash
  todo -list
  ```
- Add a new todo:
  ```bash
  <!-- normal use -->
  todo -add sample-1 task
  <!-- or receive data from another program -->
  echo sample-2 task| todo -add
  ```
- Mark a todo as completed:
  ```bash
   todo -complete <index>
  ```
- Delete a todo:
  ```bash
   todo -delete <index>
  ```
  Replace `<index>` with the task number displayed in the list.
