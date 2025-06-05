package main

import (
	"context"
	"log"
	"net"

	__tasks "github.com/petermazzocco/grpc-todo/tasks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	__tasks.UnimplementedTodoServiceServer
	tasks map[string]*__tasks.Task
}

func NewServer() *Server {
	return &Server{
		tasks: make(map[string]*__tasks.Task),
	}
}

func (s *Server) CreateTask(ctx context.Context, task *__tasks.Task) (*__tasks.Task, error) {
	log.Printf("Creating task: ID=%s, Title=%s, Description=%s, Completed=%v", task.Id, task.Title, task.Description, task.Completed)
	task.Completed = false
	s.tasks[task.Id] = task
	return task, nil
}

func (s *Server) ReadTask(ctx context.Context, req *__tasks.TaskRequest) (*__tasks.Task, error) {
	log.Printf("Reading task: ID=%s", req.Id)
	task, exists := s.tasks[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "task with ID %s not found", req.Id)
	}
	return task, nil
}

func (s *Server) UpdateTask(ctx context.Context, task *__tasks.Task) (*__tasks.Task, error) {
	log.Printf("Updating task: ID=%s, Title=%s, Description=%s", task.Id, task.Title, task.Description)
	_, exists := s.tasks[task.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "task with ID %s not found", task.Id)
	}
	s.tasks[task.Id] = task
	return task, nil
}

func (s *Server) DeleteTask(ctx context.Context, req *__tasks.TaskRequest) (*emptypb.Empty, error) {
	log.Printf("Deleting task: ID=%s", req.Id)
	_, exists := s.tasks[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "task with ID %s not found", req.Id)
	}
	delete(s.tasks, req.Id)
	return &emptypb.Empty{}, nil
}

func (s *Server) CompleteTask(ctx context.Context, req *__tasks.TaskComplete) (*__tasks.Task, error) {
	log.Printf("Completing task: ID=%s, Completed=%t", req.Id, req.Completed)
	task, exists := s.tasks[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "task with ID %s not found", req.Id)
	}
	task.Completed = req.Completed
	s.tasks[req.Id] = task
	return task, nil
}

func RunServer() {
	log.Println("Server started")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	s := NewServer()

	grpcServer := grpc.NewServer()

	__tasks.RegisterTodoServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
