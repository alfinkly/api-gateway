package handlers

import (
	"github.com/alfinkly/api-gateway/internal/application/user"
	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Register(ctx fiber.Ctx) error {
	userResp :=
		h.service.Register(ctx.Context())
}
