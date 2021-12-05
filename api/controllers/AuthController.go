package controllers

import (
	"fotongo/app/modules/auth"
	authDtos "fotongo/app/modules/auth/dtos"
	"fotongo/app/utils/baseCommands"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service auth.AuthInterface
}

func NewAuthController(service auth.AuthInterface) *AuthController {
	return &AuthController{
		service,
	}
}

func (s *AuthController) Route(f fiber.Router) {
	f.Post("/login", s.loginUserHandler)
	f.Post("/register", s.registerHandler)
}

// @Description Login in a User
// @Tags Authentication
// @Accept  json
// @Accept mpfd
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param payload body authDtos.LoginRequest true "Request Payload"
// @Success 200 {object} dtos.JSONSuccessResponses
// @Router /v1/login [POST]
func (s *AuthController) loginUserHandler(c *fiber.Ctx) error {
	request := new(authDtos.LoginRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			baseCommands.BadRequestResponse(err.Error()),
		)
	}

	errors := baseCommands.ValidateRequest(*request)
	if !reflect.ValueOf(errors.Data).IsNil() {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	result, err := s.service.Login(*request)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// @Description Register a user
// @Tags Authentication
// @Accept  json
// @Accept mpfd
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param payload body authDtos.RegisterRequest true "Request Payload"
// @Success 200 {object} dtos.JSONSuccessResponses
// @Router /v1/register [POST]
func (s *AuthController) registerHandler(c *fiber.Ctx) error {
	request := new(authDtos.RegisterRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := baseCommands.ValidateRequest(*request)
	if !reflect.ValueOf(errors.Data).IsNil() {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	result, err := s.service.Register(*request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
