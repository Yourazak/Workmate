package main

import (
	"Workmate/handlers"
	"Workmate/storage"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	store := storage.NewTaskStorage()
	handler := handlers.NewTaskHandler(store)

	r := chi.NewRouter()

	r.Post("/tasks", handler.CreateTask)
	r.Get("/tasks/{id}", handler.GetTask)
	r.Delete("/tasks/{id}", handler.DeleteTask)

	log.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
