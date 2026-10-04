package model

import "time"

type User struct {
	Fullname    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
}
