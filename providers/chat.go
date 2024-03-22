package providers

import "simplylab/model"

type ChatProvider struct {
}

func (p ChatProvider) CheckUserMessageLimitedIn30Seconds(user model.User) bool {
	return false
}

func (p ChatProvider) CheckUserMessageLimitedInDaily(user model.User) bool {
	return false
}

func (p ChatProvider) AddChatMessages(messages []model.NewMessage) int {
	return 0
}

func (p ChatProvider) GetUserChatMessages(user model.User, limit int) []model.Message {
	return make([]model.Message, 0)
}

func (p ChatProvider) GetUserChatMessagesCountToday(user model.User) int {
	return 0
}
