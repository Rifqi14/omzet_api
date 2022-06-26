package v1

import (
	"github.com/Rifqi14/omzet_api/server/http/handlers"
	"github.com/gofiber/fiber/v2"
)

type Routers struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (routers Routers) V1Route() {
	apiV1 := routers.RouteGroup.Group("/v1")

	authenticationRoute := AuthenticationRoute{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	authenticationRoute.AuthRoute()

	omzetRoute := OmzetRoute{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	omzetRoute.OmzetRoute()
}
