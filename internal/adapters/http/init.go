package http

import (
	"github.com/alfinkly/api-gateway/internal/application/user"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func InitServer() {
	server := fiber.New()

	server.Use(recover.New(recover.Config{EnableStackTrace: true}))
	server.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"*"},
		AllowOrigins:     []string{"*"},
		AllowCredentials: false,
		AllowMethods: []string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		},
	}))

	RegisterPublicRoutes(server)
}

func RegisterPublicRoutes(server *fiber.App) {
	group := server.Group("/api/public")
	repo :=
	service := user.NewService()

	group.Post("/register", )
}
