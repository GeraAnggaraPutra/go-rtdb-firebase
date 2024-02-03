package controller

import "go-firebase/src/domain/relay/service"

type RelayController struct {
	service *service.RelayService
}

func NewRelayController(service *service.RelayService) *RelayController {
	return &RelayController{service}
}
