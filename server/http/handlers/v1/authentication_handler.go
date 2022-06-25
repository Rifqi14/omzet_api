package v1

import (
	ihandlers "github.com/Rifqi14/omzet_api/domain/handlers"
	"github.com/Rifqi14/omzet_api/domain/request"
	"github.com/Rifqi14/omzet_api/package/messages"
	"github.com/Rifqi14/omzet_api/package/responses"
	"github.com/Rifqi14/omzet_api/server/http/handlers"
	v1 "github.com/Rifqi14/omzet_api/usecase/v1"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthenticationHandler struct {
	handlers.Handler
}

func NewAuthenticationHandler(handler handlers.Handler) ihandlers.IAuthenticationHandler {
	return &AuthenticationHandler{
		Handler: handler,
	}
}

func (h AuthenticationHandler) Login(ctx *fiber.Ctx) (err error) {
	req := new(request.LoginRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, fiber.StatusUnauthorized, messages.Unauthorized, err)).Send(ctx)
	}
	if err := h.Validate.Struct(req); err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, fiber.StatusUnauthorized, messages.Unauthorized, err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewAuthenticationUseCase(h.UcContract)
	res, err := uc.Login(req)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, fiber.StatusUnauthorized, messages.Unauthorized, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, "Login success")).Send(ctx)
}

func (h AuthenticationHandler) GetCurrentUser(ctx *fiber.Ctx) (err error) {
	panic("implement me")
}

func (h AuthenticationHandler) Register(ctx *fiber.Ctx) (err error) {
	panic("implement me")
}
