
# 📝 Go To-Do CLI Application

A simple yet powerful command-line interface (CLI) application built with Go to help you manage your daily tasks efficiently. Add, list, update, delete, search, and sort your to-do items directly from your terminal!

This application is powered by the [Cobra](https://github.com/spf13/cobra) library — a popular framework for building modern CLI apps in Go.

---

## 🚀 Features

- **Add Tasks**: Quickly add new to-do items with optional completion status.
- **List Tasks**: View your list with:
  - Only uncompleted tasks (default)
  - All tasks (with `--all`)
  - Limited number of tasks
  - Sorted by date (asc/desc)
- **Update Tasks**: Modify descriptions or mark them complete/pending.
- **Delete Tasks**: Remove tasks by ID.
- **Search Tasks**: Find tasks by ID or keywords.
- **Persistent Storage**: Saves tasks to `data.csv` — all changes persist across sessions.

---

## 🛠️ Installation

### Prerequisites

- [Go](https://golang.org/dl) must be installed.

### Steps

1. Clone the repository:

```bash
git clone https://github.com/YOUR_GITHUB_USERNAME/your-todo-app.git
cd your-todo-app
```

> Replace the URL with your actual repository.

2. Build the application:

```bash
go build -o todo .
```

This creates a `todo` executable (or `todo.exe` on Windows).

3. *(Optional)* Add it to your system’s `PATH`:

- **Linux/macOS:**

```bash
sudo mv todo /usr/local/bin/
```

- **Windows:**

Move `todo.exe` to a folder already in your system’s `Path`, e.g., `C:\Windows`.

---

## 💡 Usage

All commands follow this structure:

```bash
todo [command] [flags]
```

---

### ✅ `todo add`

Add a new task.

```bash
todo add "Buy groceries"
todo add "Submit report" --done Yes
```

**Flags**:
- `-d, --done <Yes|No>`: Set initial completion status. Default: `No`.

---

### 📋 `todo list`

List tasks.

```bash
todo list
todo list --all
todo list --listcount 3
todo list --all --listcount 5 --descending
```

**Flags**:
- `-a, --all`: Show completed tasks.
- `-c, --listcount <number>`: Limit displayed tasks.
- `-d, --descending`: Sort by date (newest first).

---

### ✏️ `todo update`

Update a task by ID.

```bash
todo update --id 1 --done Yes
todo update --id 2 --task "Call John"
todo update --id 3 --task "Review docs" --done No
```

**Required Flags**:
- `-i, --id <number>`: Task ID.
- `-t, --task <string>`: New task description.
- `-d, --done <Yes|No>`: Update status.

---

### ❌ `todo delete`

Delete a task by ID.

```bash
todo delete 4
```

---

### 🔍 `todo search`

Search by ID or keyword.

```bash
todo search --id 5
todo search --task "report"
```

**Flags**:
- `-i, --id <number>`
- `-t, --task <string>`

> You must use either `--id` or `--task`, not both.

---

## 📁 Project Structure

```
your-todo-app/
├── main.go                 # Entry point
├── cmd/                    # Cobra commands
│   ├── root.go
│   ├── add.go
│   ├── list.go
│   ├── update.go
│   ├── delete.go
│   └── search.go
├── handlers/               # CSV file logic
│   ├── add.go
│   ├── list.go
│   ├── update.go
│   ├── delete.go
│   └── search.go
├── utils/
│   └── utils.go            # Helper functions
└── data.csv                # Task storage
```

---

## 🤝 Contributing

Contributions are welcome! Please open an issue or pull request with improvements, ideas, or bug fixes.

---

## 📄 License

This project is licensed under the **MIT License** — see the [LICENSE](./LICENSE) file for details.

---

## 📧 Contact

Created by **[AliYahyavy]** — feel free to reach out!  
Email : **[aliyahyavi77@yahoo.com]**  
Telegram ID : **[@ali_yahyavii]**
