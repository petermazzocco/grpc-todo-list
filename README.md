# grpc-todo

A simple todo list application built with Go and gRPC, featuring three different interfaces: gRPC server, CLI client, and HTTP API.

## Features

- **gRPC Server**: Core todo service running on port 9000
- **CLI Client**: Command-line interface for task management
- **HTTP API**: RESTful API server running on port 3000
- Basic CRUD operations: Create, Read, Update, Delete, and Complete tasks

## Usage

### Starting the gRPC Server
```bash
go run . -mode=server
```

### Using the CLI Client
```bash
# Create a new task
go run . -mode=cli -action=new -id="1" -title="Buy groceries" -desc="Milk, bread, eggs"

# Get a task
go run . -mode=cli -action=get -id="1"

# Update a task
go run . -mode=cli -action=update -id="1" -title="Updated title" -desc="Updated description"

# Mark task as complete
go run . -mode=cli -action=done -id="1"

# Delete a task
go run . -mode=cli -action=delete -id="1"
```

### Starting the HTTP API Server
```bash
go run . -mode=api
```

#### API Endpoints
- `POST /tasks/new` - Create a new task
- `GET /tasks/{id}` - Get a task by ID
- `PUT /tasks/{id}` - Update a task
- `POST /tasks/{id}/done` - Mark task as complete
- `DELETE /tasks/{id}` - Delete a task

## Architecture

- `main.go` - Entry point with mode selection and CLI argument parsing
- `server.go` - gRPC server implementation with in-memory task storage
- `cli.go` - CLI client functions that communicate with the gRPC server
- `api.go` - HTTP API server using Chi router that proxies to gRPC functions
- `tasks/` - Generated protobuf files for gRPC communication
