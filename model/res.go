package model

type GetAiChatResponseOutput struct {
	Response string `json:"response"`
}

type GetChatStatusTodayOutput struct {
	UserName UserName `json:"user_name"`
	ChatCnt  int   `json:"chat_cnt"`
}

type UserChatMessage struct {
	Type MessageRoleType `json:"type"`
	Text string          `json:"text"`
}

type GetUserChatHistoryOutput []UserChatMessage
