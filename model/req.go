package model

type GetAiChatResponseInput struct {
	Message  string   `json:"message"`
	UserName UserName `json:"user_name"`
}
