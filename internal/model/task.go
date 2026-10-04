package model

import "time"

type Task struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}
