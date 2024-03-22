package services

import (
	"simplylab/model"
	"simplylab/providers"
)

type Services struct {
	ctx *model.ServiceContext
	pvd *providers.Providers
}

func NewServices(ctx *model.ServiceContext, pvd *providers.Providers) *Services {
	return &Services{
		ctx: ctx,
		pvd: pvd,
	}
}

func (s Services) Ping() PingService {
	return PingService{ctx: s.ctx, pvd: s.pvd}
}

func (s Services) Chat() ChatService {
	return ChatService{ctx: s.ctx, pvd: s.pvd}
}
