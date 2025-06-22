package storage

import (
	"Workmate/task"
	"sync"
)

type TaskStorage struct {
	mu    sync.RWMutex
	tasks map[string]*task.Task
}

func NewTaskStorage() *TaskStorage {
	return &TaskStorage{
		tasks: make(map[string]*task.Task),
	}
}
func (s *TaskStorage) Add(t *task.Task) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tasks[t.ID] = t
}

func (s *TaskStorage) Get(id string) (*task.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.tasks[id]
	return t, ok
}

func (s *TaskStorage) Delete(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.tasks, id)
}
