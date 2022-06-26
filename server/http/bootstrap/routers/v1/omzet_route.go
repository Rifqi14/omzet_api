package v1

import (
	"net/http"

	"github.com/Rifqi14/omzet_api/server/http/handlers"
	v1 "github.com/Rifqi14/omzet_api/server/http/handlers/v1"
	"github.com/Rifqi14/omzet_api/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

type OmzetRoute struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (r OmzetRoute) OmzetRoute() {
	handler := v1.NewOmzetHandler(r.Handler)
	jwtMiddleware := middlewares.JwtMiddleware{Contract: r.Handler.UcContract}

	omzetRoute := r.RouteGroup.Group("/omzet")
	omzetRoute.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON("Api is healthy")
	})
	omzetRoute.Use(jwtMiddleware.New)
	omzetRoute.Get("/merchant/:id", handler.GetReportMerchant)
}
