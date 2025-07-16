# Quiz App in Go

This is a simple command-line quiz application written in Go. It reads questions and answers from a CSV file and quizzes the user, with an optional time limit for the entire quiz.

---

## Features

- Reads quiz questions and answers from a CSV file.
- Configurable time limit for the quiz (default: 30 seconds).
- Displays the final score at the end or when time runs out.
- Command-line flags for easy configuration.

---

## Getting Started

### Prerequisites

- [Go 1.18+](https://golang.org/dl/)

### Installation

1. **Clone the repository:**
   ```sh
   git clone <your-repo-url>
   cd goexcercises/quizApp
   ```

2. **Place your `problems.csv` file in the same directory.**
   - The CSV file should have the format:
     ```
     question,answer
     5+5,10
     7+3,10
     1+1,2
     ```

---

## Usage

### Run the Quiz

```sh
go run main.go
```

- By default, it will look for `problems.csv` in the current directory and use a 30-second timer.

### Command-Line Flags

- **Specify a different CSV file:**
  ```sh
  go run main.go -csv=yourfile.csv
  ```
- **Set a custom time limit (in seconds):**
  ```sh
  go run main.go -timelimit=60
  ```
- **Both together:**
  ```sh
  go run main.go -csv=yourfile.csv -timelimit=60
  ```

---

## Example

```
Question #1: 5+5=
10
Question #2: 7+3=
10
...
Your final score is: 2/2
```

If the timer runs out, the quiz ends and your score is displayed.

---

## License

MIT

---
