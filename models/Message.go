package models

import "time"

type Message struct {
	ID      int       `json:"id" db:"id"`
	Message string    `json:"message" db:"messsage"`
	Date    time.Time `json:"date" db:"date"`

	WorkspaceID string `json:"workspace_id" db:"workspace_id"`
	Username    string `json:"username" db:"username"`
}
