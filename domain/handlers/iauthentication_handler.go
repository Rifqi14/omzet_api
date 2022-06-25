package handlers

import "github.com/gofiber/fiber/v2"

type IAuthenticationHandler interface {
	Login(ctx *fiber.Ctx) (err error)

	GetCurrentUser(ctx *fiber.Ctx) (err error)

	Register(ctx *fiber.Ctx) (err error)
}
