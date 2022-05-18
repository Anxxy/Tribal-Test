package services

import (
	"gopkg.in/resty.v1"
	"tribal/services/interfaces"
)

type service struct {
	RestClient *resty.Client
}

func NewService(restClient *resty.Client) interfaces.Service {
	return service{
		RestClient: restClient,
	}
}
