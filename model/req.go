package model

type GetAiChatResponseInput struct {
	Message  string   `json:"message"`
	UserName UserName `json:"user_name"`
}

type GetUserChatHistoryInput struct {
	UserName UserName `json:"user_name"`
	LastN    int      `json:"last_n"`
}

type GetChatStatusTodayInput struct {
	UserName UserName `json:"user_name"`
}

type NewMessage struct {
	UserId string          `json:"user_id"`
	Type   MessageRoleType `json:"type"`
	Text   string          `json:"text"`
}
