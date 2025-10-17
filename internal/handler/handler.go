package handler

import (
	authv1 "github.com/versegeek/go-generated/gen/proto/auth/v1"
	oauth2V1 "github.com/versegeek/go-generated/gen/rest/oauth2/v1"
	"github.com/versegeek/go-skeleton/internal/service"
)

// var _ Handler = (*handler)(nil)

type (
	Handler interface {
		OAuth2Handler
		VersionHandler
		authv1.ClientAPIServer
		oauth2V1.ServerInterface
	}

	handler struct {
		service service.IService
		authv1.UnimplementedClientAPIServer
		oauth2V1.ServerInterfaceWrapper
	}
)

func New(service service.IService) Handler {
	return &handler{
		service: service,
	}
}
