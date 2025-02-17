package handler

import (
	"github.com/alfisar/jastip-import/domain"

	"github.com/gofiber/fiber/v2"
)

func HandlerRegistration(c *fiber.Ctx) (domain.User, error) {
	request := domain.User{}
	errData := c.BodyParser(&request)
	if errData != nil {
		return request, errData
	}

	return request, nil
}

func HandlerVerify(c *fiber.Ctx) (domain.UserVerifyOtpRequest, error) {
	request := domain.UserVerifyOtpRequest{}
	errData := c.BodyParser(&request)
	if errData != nil {
		return request, errData
	}

	return request, nil
}

func HandlerResend(c *fiber.Ctx) (domain.UserResendOtpRequest, error) {
	request := domain.UserResendOtpRequest{}
	errData := c.BodyParser(&request)
	if errData != nil {
		return request, errData
	}

	return request, nil
}

func HandlerLogin(c *fiber.Ctx) (domain.UserLoginRequest, error) {
	request := domain.UserLoginRequest{}
	errData := c.BodyParser(&request)
	if errData != nil {
		return request, errData
	}

	return request, nil
}

func HandlerUpdateProfile(c *fiber.Ctx) (map[string]any, error) {
	request := map[string]any{}
	errData := c.BodyParser(&request)
	if errData != nil {
		return request, errData
	}

	return request, nil
}

func HandlerpostAddress(c *fiber.Ctx) (map[string]any, error) {
	request := map[string]any{}
	errData := c.BodyParser(&request)
	if errData != nil {
		return request, errData
	}

	return request, nil
}

func HandlerPostSchedule(c *fiber.Ctx) (domain.TravelSchRequest, error) {
	request := domain.TravelSchRequest{}
	errData := c.BodyParser(&request)
	if errData != nil {
		return request, errData
	}

	return request, nil
}
