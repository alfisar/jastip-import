package handler

import (
	"fmt"
	"strconv"

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

func HandlerPostProducts(c *fiber.Ctx) (domain.ProductData, error) {
	request := domain.ProductData{}
	errData := c.BodyParser(&request)
	if errData != nil {
		return request, errData
	}

	return request, nil
}

func HandlerParamSch(c *fiber.Ctx) (domain.Params, error) {
	errMessage := ""

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		errMessage = "page tidak valid"
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		if errMessage != "" {
			errMessage += ", limit tidak valid"
		} else {
			errMessage = "limit tidak valid"
		}

	}

	status, err := strconv.Atoi(c.Query("status"))
	if err != nil {
		if errMessage != "" {
			errMessage += ", status tidak valid"
		} else {
			errMessage = "status tidak valid"
		}

	}

	if errMessage != "" {
		return domain.Params{}, fmt.Errorf(errMessage)
	}

	return domain.Params{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
		Status: status,
	}, nil
}

func HandlerParamCountries(c *fiber.Ctx) (domain.Params, error) {
	errMessage := ""

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		errMessage = "page tidak valid"
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		if errMessage != "" {
			errMessage += ", limit tidak valid"
		} else {
			errMessage = "limit tidak valid"
		}

	}

	if errMessage != "" {
		return domain.Params{}, fmt.Errorf(errMessage)
	}

	return domain.Params{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
	}, nil
}

func HandlerParamProducts(c *fiber.Ctx) (domain.Params, error) {
	errMessage := ""

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		errMessage = "page tidak valid"
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		if errMessage != "" {
			errMessage += ", limit tidak valid"
		} else {
			errMessage = "limit tidak valid"
		}

	}

	status, err := strconv.Atoi(c.Query("status"))
	if err != nil {
		if errMessage != "" {
			errMessage += ", status tidak valid"
		} else {
			errMessage = "status tidak valid"
		}

	}

	if errMessage != "" {
		return domain.Params{}, fmt.Errorf(errMessage)
	}

	return domain.Params{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
		Status: status,
	}, nil
}

func HandlerUpdate(c *fiber.Ctx) (map[string]any, error) {
	request := map[string]any{}
	errData := c.BodyParser(&request)
	if errData != nil {
		return request, errData
	}

	return request, nil
}

func HandlerPathID(c *fiber.Ctx) (int, error) {
	errMessage := ""

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		errMessage = "id tidak valid"
		return 0, fmt.Errorf(errMessage)
	}

	return id, nil
}
