package handler

import (
	"github.com/amartha-test/usecase"
)

type Server struct {
	UseCase usecase.UseCaseInterface
}

type NewServerOptions struct {
	UseCase usecase.UseCaseInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		UseCase: opts.UseCase,
	}
}
