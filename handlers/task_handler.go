package handlers

import (
	"Workmate/storage"
	"Workmate/task"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type TaskHandler struct {
	Store *storage.TaskStorage
}

func NewTaskHandler(store *storage.TaskStorage) *TaskHandler {
	return &TaskHandler{Store: store}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	id := uuid.New().String()
	t := &task.Task{
		ID:        id,
		Status:    task.StatusPending,
		CreatedAt: time.Now(),
		Duration:  "0s",
	}

	h.Store.Add(t)

	go func(t *task.Task) {
		t.Status = task.StatusInProgress
		start := time.Now()
		time.Sleep(3 * time.Minute)
		t.Duration = time.Since(start).String()
		t.Status = task.StatusDone
	}(t)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	t, ok := h.Store.Get(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	h.Store.Delete(id)
	w.WriteHeader(http.StatusNoContent)
}
