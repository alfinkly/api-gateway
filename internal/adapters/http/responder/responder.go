package respond

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v3"
	"log"
	"reflect"
)

type Formatter[T any] func(T) any

type Config[T any] struct {
	Status      int
	Formatter   Formatter[T]
	SuppressNil bool
}

func Respond[T any](c fiber.Ctx, data T, err error, opts ...Config[T]) error {
	cfg := defaultConfig[T]()
	if len(opts) > 0 {
		cfg = opts[0]
	}

	log.Printf("[%s] %s", c.Method(), c.OriginalURL())

	if err != nil {
		log.Printf("[ERROR] %v", err)

		switch {
		case errors.Is(err, sql.ErrNoRows):
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "resource not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	if isNil(data) {
		if cfg.SuppressNil {
			return c.SendStatus(cfg.Status)
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "resource is nil",
		})
	}

	var out any = data
	if cfg.Formatter != nil {
		out = cfg.Formatter(data)
	}

	return c.Status(cfg.Status).JSON(out)
}

func defaultConfig[T any]() Config[T] {
	return Config[T]{
		Status:      fiber.StatusOK,
		SuppressNil: false,
	}
}

func isNil[T any](v T) bool {
	val := reflect.ValueOf(v)
	if !val.IsValid() {
		return true
	}
	if val.Kind() == reflect.Ptr || val.Kind() == reflect.Slice || val.Kind() == reflect.Map || val.Kind() == reflect.Interface {
		return val.IsNil()
	}
	return false
}
