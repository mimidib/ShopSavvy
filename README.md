# My Go Project

This is a simple Go application that demonstrates how to create a basic "Hello, World!" program.

## Project Structure

```
my-go-project
├── cmd
│   └── main.go       # Entry point of the application
├── pkg
│   └── hello
│       └── hello.go  # Contains the SayHello function
└── README.md         # Project documentation
```

## Getting Started

To run this application, you need to have Go installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).

### Running the Application

1. Clone the repository or download the project files.
2. Navigate to the project directory:
   ```
   cd my-go-project
   ```
3. Run the application using the following command:
   ```
   go run cmd/main.go
   ```

### Purpose

This project serves as a basic introduction to Go programming, showcasing how to structure a Go application with separate packages and an entry point. The application simply prints "Hello, World!" to the console.