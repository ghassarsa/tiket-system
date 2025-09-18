package config

import (
	"reflect"

	"github.com/TesyarRAz/ticket-system/pkg/exception"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewFiber(cfg Config) *fiber.App {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: false,
	})

	return app
}

var httpErrorCodeMap = map[reflect.Type]int{
	reflect.TypeOf(&exception.NotFound{}):     fiber.StatusNotFound,
	reflect.TypeOf(&exception.Internal{}):     fiber.StatusInternalServerError,
	reflect.TypeOf(&exception.BadRequest{}):   fiber.StatusBadRequest,
	reflect.TypeOf(&exception.Unauthorized{}): fiber.StatusUnauthorized,
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	ers := make(map[string]any)
	msg := err.Error()
	code := fiber.StatusInternalServerError

	response := map[string]any{
		"message": msg,
		"errors":  ers,
	}

	switch e := err.(type) {
	case *fiber.Error:
		code = e.Code
		msg = e.Message
	case validator.ValidationErrors:
		msg = "Validation error"
		for _, fe := range e {
			ers[fe.Field()] = fiber.Map{
				"tag": fe.Tag(),
				"val": fe.Param(),
			}
		}
		code = fiber.StatusUnprocessableEntity
	default:
		t := reflect.TypeOf(err)
		if codeMapped, ok := httpErrorCodeMap[t]; ok {
			code = codeMapped
			break
		}
	}

	return ctx.Status(code).JSON(response)
}
