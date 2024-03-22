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
