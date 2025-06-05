package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	__tasks "github.com/petermazzocco/grpc-todo/tasks"
)

func MarshalAnEncodeJSON(w http.ResponseWriter, v any) {
	data, err := json.Marshal(v)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(string(data))
}

func StartHTTPServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/tasks", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				id := chi.URLParam(r, "id")
				task := GetTask(id)
				MarshalAnEncodeJSON(w, task)
			})
			r.Post("/done", func(w http.ResponseWriter, r *http.Request) {
				id := chi.URLParam(r, "id")
				task := MarkComplete(id)
				MarshalAnEncodeJSON(w, task)
			})
			r.Put("/", func(w http.ResponseWriter, r *http.Request) {
				id := chi.URLParam(r, "id")

				var t __tasks.Task
				if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
					http.Error(w, "Invalid JSON", http.StatusBadRequest)
					return
				}

				task := UpdateTask(id, t.Title, t.Description)
				MarshalAnEncodeJSON(w, task)
			})
			r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
				id := chi.URLParam(r, "id")
				msg := DeleteTask(id)
				MarshalAnEncodeJSON(w, msg)
			})
		})
		r.Post("/new", func(w http.ResponseWriter, r *http.Request) {
			var t __tasks.Task
			if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			task := CreateTask(t.Id, t.Title, t.Description)
			MarshalAnEncodeJSON(w, task)
		})
	})
	log.Println("API server started on port :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
