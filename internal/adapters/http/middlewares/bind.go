package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validate = validator.New()

func BindAndValidate[T any]() fiber.Handler {
	return func(c fiber.Ctx) error {
		var body T

		if err := c.Bind().Body(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "invalid JSON",
				"details": err.Error(),
			})
		}

		if err := validate.Struct(body); err != nil {
			errs := make(map[string]string)
			for _, e := range err.(validator.ValidationErrors) {
				errs[e.Field()] = e.ActualTag()
			}

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "validation failed",
				"details": errs,
			})
		}

		c.Locals("body", body)
		return c.Next()
	}
}
