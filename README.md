# Go CLI To-Do App

A simple command-line To-Do list application written in Go. It allows you to add, list, mark as done, and delete tasks directly from your terminal. Tasks are stored in a local `todo.json` file for persistence.

---

## Features

- Add new tasks
- List all tasks with their current status
- Mark tasks as completed
- Delete tasks
- Task data is saved to a `todo.json` file

---


## Usage

Run the app with the following command structure:

```bash
./todo <command> [arguments]
```

### Available Commands

| Command          | Description                       | Example                        |
|------------------|-----------------------------------|--------------------------------|
| `add`            | Add a new task                    | `./todo add "Learn Go"`        |
| `list`           | List all tasks                    | `./todo list`                  |
| `done <taskID>`  | Mark a task as completed          | `./todo done 1`                |
| `delete <taskID>`| Delete a task                     | `./todo delete 2`              |

---

## Data Persistence

Tasks are saved to a local `todo.json` file in the project directory. This ensures that tasks persist across sessions.

---

## Example

### 1. Add a Task

```bash
$ ./todo add "Write Go CLI app"
Added task: Write Go CLI app
```

### 2. List All Tasks

```bash
$ ./todo list
1. [❌] Write Go CLI app
```

### 3. Mark Task as Done

```bash
$ ./todo done 1
Marked task done: Write Go CLI app
```

### 4. List All Tasks Again

```bash
$ ./todo list
1. [✅] Write Go CLI app
```

### 5. Delete Task

```bash
$ ./todo delete 1
Deleted task: Write Go CLI app
```

---


## Potential Improvements

- Add support for due dates or task priorities
- Filter tasks (e.g., list only completed or pending tasks)
- Task editing feature (update description)
- Use a more advanced CLI library like Cobra or urfave/cli for improved flexibility
- Add colorful output using ANSI codes for better visibility
