package services

import (
	"simplylab/model"
	"simplylab/providers"
)

type PingService struct {
	ctx *model.ServiceContext
	pvd *providers.Providers
}

