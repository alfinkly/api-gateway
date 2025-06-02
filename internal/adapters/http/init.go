package http

import (
	"github.com/alfinkly/api-gateway/internal/adapters/http/handlers"
	middleware "github.com/alfinkly/api-gateway/internal/adapters/http/middlewares"
	"github.com/alfinkly/api-gateway/internal/application/user"
	"github.com/alfinkly/api-gateway/internal/domain"
	"github.com/alfinkly/api-gateway/internal/repository"
	"github.com/alfinkly/api-gateway/internal/repository/sqlc"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func StartServer(queries *sqlc.Queries, port string) (err error) {
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

	RegisterPublicRoutes(server, queries)

	return server.Listen(port)
}

func RegisterPublicRoutes(server *fiber.App, queries *sqlc.Queries) {
	group := server.Group("/api/public")
	r := repository.NewUserRepo(queries)
	s := user.NewService(r)
	h := handlers.NewUserHandler(s)

	group.Post("/register", h.Register, middleware.BindAndValidate[domain.RegisterResponse]())
}
