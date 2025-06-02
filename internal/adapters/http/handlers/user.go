package handlers

import (
	respond "github.com/alfinkly/api-gateway/internal/adapters/http/responder"
	"github.com/alfinkly/api-gateway/internal/application/user"
	"github.com/alfinkly/api-gateway/internal/domain"
	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	service *user.Service
}

func NewUserHandler(service *user.Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Register(ctx fiber.Ctx) error {
	userResp := ctx.Locals("body").(domain.RegisterResponse)
	out, err := h.service.Register(ctx.Context(), userResp)
	return respond.Created(ctx, out, err)
}
