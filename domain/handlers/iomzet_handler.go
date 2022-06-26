package handlers

import "github.com/gofiber/fiber/v2"

type IOmzetHandler interface {
	GetReportMerchant(ctx *fiber.Ctx) (err error)

	GetReportOutlet(ctx *fiber.Ctx) (err error)
}
