package task

import "time"

type Status string

const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

type Task struct {
	ID        string    `json:"id"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	Duration  string    `json:"duration"`
}
