package handler

import (
	"INi-Wallet/service"
	"INi-Wallet/usecase"
)

type Handler struct {
	userUseCase usecase.UserUseCase
	jwtService  service.JWTService
}

type HandlerConfig struct {
	UserUseCase usecase.UserUseCase
	JWTService  service.JWTService
}

func NewHandler(c *HandlerConfig) *Handler {
	return &Handler{
		userUseCase: c.UserUseCase,
		jwtService:  c.JWTService,
	}
}
