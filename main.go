package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	// Flags for CLI
	mode := flag.String("mode", "server", "Mode to run: server or client")
	action := flag.String("action", "action", "Create, update, delete a task")

	id := flag.String("id", "id", "The ID of the new task")
	title := flag.String("title", "title", "The title of the new task")
	desc := flag.String("desc", "description", "The description of the task")
	flag.Parse()

	// Choose with mode you want the client to be (CLI or API)
	switch *mode {
	case "cli":
		switch *action {
		case "get":
			task := GetTask(*id)
			log.Println(task)
		case "new":
			task := CreateTask(*id, *title, *desc)
			log.Println(task)
		case "done":
			task := MarkComplete(*id)
			log.Println(task)
		case "update":
			task := UpdateTask(*id, *title, *desc)
			log.Println(task)
		case "delete":
			msg := DeleteTask(*id)
			log.Println(msg)
		}
	case "api":
		StartHTTPServer()
	case "server":
		RunServer()
	default:
		log.Fatalf("Invalid mode: %s. Use 'server', 'api' or 'cli'", *mode)
		os.Exit(1)
	}
}
