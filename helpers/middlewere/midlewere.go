package middlewere

import (
	"errors"

	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/consts"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"github.com/alfisar/jastip-import/helpers/helper"
	"github.com/alfisar/jastip-import/helpers/jwthandler"
	"github.com/alfisar/jastip-import/helpers/response"
	repository "github.com/alfisar/jastip-import/repository/redis"
	"github.com/gofiber/fiber/v2"
)

// AuthenticateMiddleware is a middleware for user authentication
type AuthenticateMiddleware struct {
	jwt *jwthandler.JwtHandler
}

func ValidationPath[T any](parse func(c *fiber.Ctx) (T, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		path, err := parse(c)
		if err != nil {
			log.Printf("Error parsing data on middleware : %s", err.Error())
			err := errorhandler.ErrValidation(err)
			response.WriteResponse(c, response.Response{}, err, fiber.StatusBadRequest)
			return nil
		}

		c.Locals("path", path)
		return c.Next()
	}
}

func Validation[T any](parse func(c *fiber.Ctx) (T, error), validate func(T) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		request, err := parse(c)
		if err != nil {
			log.Printf("Error parsing data on middleware : %s", err.Error())
			err := errorhandler.ErrValidation(err)
			response.WriteResponse(c, response.Response{}, err, fiber.StatusBadRequest)
			return nil
		}

		if validate != nil {
			err = validate(request)
			if err != nil {
				log.Printf("Error validation data on middleware : %s", err.Error())
				err := errorhandler.ErrValidation(err)
				response.WriteResponse(c, response.Response{}, err, fiber.StatusBadRequest)
				return nil
			}
		}

		c.Locals("validatedData", request)
		return c.Next()
	}
}

// NewAuthenticateMiddleware create objcet of authenticate middleware
func NewAuthenticateMiddleware(jwt *jwthandler.JwtHandler) *AuthenticateMiddleware {
	return &AuthenticateMiddleware{
		jwt: jwt,
	}
}

// Authenticate authenticates the user who accessed the handler
func (obj *AuthenticateMiddleware) Authenticate(ctx *fiber.Ctx) error {

	ctx.Locals("time", time.Now())
	if ctx.Method() == "OPTIONS" {
		return ctx.Next()
	}

	poolData := domain.DataPool.Get().(*domain.Config)

	tokenString, errData := getTokenRequest(ctx)
	if errData != nil {
		err := errorhandler.ErrMiddleware(errorhandler.ErrCodeInvalidAuth, errorhandler.ErrMsgTokenInvalid, errData)

		return ctx.Status(http.StatusUnauthorized).JSON(err)
	}

	token, errData := helper.DecryptAES256CBC(poolData.Hash.Key, tokenString)
	if errData != nil {
		log.Printf("Error hashing aes 256 token on func login : %s", errData.Error())

		err := errorhandler.ErrHashing(errData)
		return ctx.Status(http.StatusUnauthorized).JSON(err)
	}

	_, claim, errData := obj.jwt.ValidationToken(token)
	if errData != nil {
		err := errorhandler.ErrMiddleware(errorhandler.ErrCodeInvalidAuth, errorhandler.ErrMsgTokenInvalid, errData)

		return ctx.Status(http.StatusUnauthorized).JSON(err)
	}

	keys := "TOKEN_" + strconv.Itoa(int(claim["user_id"].(float64)))
	repoRedis := repository.NewRedisRepository()
	result, errData := repoRedis.Get(ctx.Context(), poolData.DBRedis[consts.RedisToken], keys)
	if errData != nil {
		err := errorhandler.ErrMiddleware(errorhandler.ErrCodeInvalidAuth, errorhandler.ErrMsgTokenInvalid, errData)

		return ctx.Status(http.StatusUnauthorized).JSON(err)
	}

	if result != tokenString || result == "" {
		if errData != nil {
			err := errorhandler.ErrMiddleware(errorhandler.ErrCodeInvalidAuth, errorhandler.ErrMsgTokenInvalid, errData)

			return ctx.Status(http.StatusUnauthorized).JSON(err)
		}
	}

	ctx.Locals("data", claim["user_id"])

	return ctx.Next()

}

func getTokenRequest(ctx *fiber.Ctx) (tokenString string, err error) {
	const bearerSchema = "Bearer "
	authHeader := ctx.Get("Authorization")
	if len(authHeader) < len(bearerSchema) {
		err = errors.New("Invalid Authorization")
		return
	}

	tokenString = authHeader[len(bearerSchema):]
	return
}
