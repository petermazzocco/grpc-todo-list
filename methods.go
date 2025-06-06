package main

import (
	"context"
	"fmt"
	"log"

	__tasks "github.com/petermazzocco/grpc-todo/tasks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient() (__tasks.TodoServiceClient, *grpc.ClientConn, error) {
	creds := insecure.NewCredentials()
	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, nil, fmt.Errorf("error occurred connecting: %s", err)
	}

	c := __tasks.NewTodoServiceClient(conn)
	return c, conn, nil
}

func GetTask(id string) *__tasks.Task {
	c, conn, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	res, err := c.ReadTask(context.Background(), &__tasks.TaskRequest{
		Id: id,
	})
	if err != nil {
		log.Fatalf("Error when calling GetTask: %s", err)
	}

	return res
}

func CreateTask(id, title, description string) *__tasks.Task {
	c, conn, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	res, err := c.CreateTask(context.Background(), &__tasks.Task{
		Id:          id,
		Title:       title,
		Description: description,
		Completed:   false,
	})
	if err != nil {
		log.Fatalf("Error when calling CreateTask: %s", err)
	}
	return res
}

func UpdateTask(id, title, description string) *__tasks.Task {
	c, conn, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	res, err := c.UpdateTask(context.Background(), &__tasks.Task{
		Id:          id,
		Title:       title,
		Description: description,
		Completed:   false,
	})
	if err != nil {
		log.Fatalf("Error when calling UpdateTask: %s", err)
	}

	return res
}

func DeleteTask(id string) string {
	c, conn, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	res, err := c.DeleteTask(context.Background(), &__tasks.TaskRequest{
		Id: id,
	})
	if err != nil {
		log.Fatalf("Error when calling DeleteTask: %s", err)
	}

	return res.String()
}

func MarkComplete(id string) *__tasks.Task {
	c, conn, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	res, err := c.CompleteTask(context.Background(), &__tasks.TaskComplete{
		Id:        id,
		Completed: true,
	})
	if err != nil {
		log.Fatalf("Error when calling MarkComplete: %s", err)
	}

	return res
}
