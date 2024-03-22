package model

import "time"

type UserID string
type UserName string
type CreatedAt time.Time
type UpdatedAt time.Time
type CreatedBy string
type UpdatedBy string

type User struct {
	ID        UserID    `json:"id"`
	Name      UserName  `json:"name"`
	CreatedAt CreatedAt `json:"created_at"`
	UpdatedAt UpdatedAt `json:"updated_at"`
}

type MessageRoleType string

const (
	MessageRoleUser MessageRoleType = "user"
	MessageRoleAI   MessageRoleType = "ai"
)

type MessageID string

type Message struct {
	ID        MessageID       `json:"id"`
	UserID    UserID          `json:"user_id"`
	Type      MessageRoleType `json:"type"`
	Text      string          `json:"text"`
	CreatedAt CreatedAt       `json:"created_at"`
	CreatedBy CreatedBy       `json:"created_by"`
	UpdatedAt UpdatedAt       `json:"updated_at"`
	UpdatedBy UpdatedBy       `json:"updated_by"`
}

type ServiceContext struct {
	User User `json:"user"`
}
