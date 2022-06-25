package v1

import (
	"github.com/Rifqi14/omzet_api/server/http/handlers"
	v1 "github.com/Rifqi14/omzet_api/server/http/handlers/v1"
	"github.com/gofiber/fiber/v2"
)

type AuthenticationRoute struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (r AuthenticationRoute) AuthRoute() {
	handler := v1.NewAuthenticationHandler(r.Handler)
	// jwtMiddleware := middlewares.JwtMiddleware{Contract: r.Handler.UcContract}

	loginRoute := r.RouteGroup.Group("/login")
	loginRoute.Post("/", handler.Login)
}
