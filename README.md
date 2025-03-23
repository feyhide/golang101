# GoLang Learning Repository

Welcome to my **GoLang Learning Repository**! 🚀

This repository documents my journey of learning **Go (Golang)** from scratch. Each folder contains code and examples related to different Go concepts, from fundamental syntax to advanced topics.

## ▶️ Running the Code

To execute a specific lesson (excluding the web application), use:

```sh
go run <foldername>/main.go
```

For example, to run the **pointers** lesson:

```sh
go run 15_pointers/main.go
```

For the WebApp, use the following command (providing the path to the configuration file):

```sh
go run cmd/webapp/main.go -config config/local.yaml
```

## 📖 Topics Covered

- **Basics** (Variables, Constants, Data Types)
- **Control Structures** (Loops, Conditionals, Switch)
- **Arrays & Slices**
- **Maps & Ranges**
- **Functions & Closures**
- **Pointers**
- **Structs & Interfaces**
- **Concurrency** (Goroutines, WaitGroups, Channels, Mutex)
- **File Handling**
- **Streaming**
- **Packages & Modules**

## 🌐 WebApp API Endpoints

The WebApp provides the following API endpoints:

- **Create Student** → `POST /api/v1/students`
- **Get Student by ID** → `GET /api/v1/students/{id}`
- **List All Students** → `GET /api/v1/students`
- **Update Student Details** → `PATCH /api/v1/students`
- **Delete Student** → `DELETE /api/v1/students/{id}`
