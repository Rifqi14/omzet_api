package handlers

import "github.com/gofiber/fiber/v2"

type IOmzetHandler interface {
	GetReport(ctx *fiber.Ctx) (err error)

	GetReportWithOutlet(ctx *fiber.Ctx) (err error)
}
