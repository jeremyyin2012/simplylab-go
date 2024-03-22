package services

import (
	"simplylab/model"
	"simplylab/providers"
)

type ChatService struct {
	ctx *model.ServiceContext
	pvd *providers.Providers
}

func (s ChatService) GetAIChatResponse(req model.GetAiChatResponseInput) model.GetAiChatResponseOutput {
	return model.GetAiChatResponseOutput{Response: "todo"}
}

func (s ChatService) GetUserChatHistory(lastN int) model.GetUserChatHistoryOutput {
	return make(model.GetUserChatHistoryOutput, 0)
}

func (s ChatService) GetChatStatusToday() model.GetChatStatusTodayOutput {
	return model.GetChatStatusTodayOutput{
		UserName: s.ctx.User.Name,
		ChatCnt:  0,
	}
}
