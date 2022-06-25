package bootstrap

import (
	"net/http"

	v1 "github.com/Rifqi14/omzet_api/server/http/bootstrap/routers/v1"
	"github.com/Rifqi14/omzet_api/server/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func (boot Bootstrap) AppRoute() {
	handlerType := handlers.Handler{
		App:           boot.App,
		UcContract:    &boot.UcContract,
		DB:            boot.DB,
		Validate:      boot.Validator,
		Translator:    boot.Translator,
		JweCredential: boot.UcContract.JweCredential,
		JwtCredential: boot.UcContract.JwtCredential,
	}

	// Route check health

	rootParentGroup := boot.App.Group("/api")
	rootParentGroup.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON("Api is healthy")
	})

	v1Routes := v1.Routers{
		RouteGroup: rootParentGroup,
		Handler:    handlerType,
	}
	v1Routes.V1Route()
}
