package routes

import (
	"crud-echo/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/unrolled/secure"
)

func Routes() *echo.Echo {

	e := echo.New()

	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:            []string{"localhost:9094", "www.google.com"},
		FrameDeny:               true,
		CustomFrameOptionsValue: "SAMEORIGIN",
		ContentTypeNosniff:      true,
		BrowserXssFilter:        true,
	})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.Use(echo.WrapMiddleware(secureMiddleware.Handler))

	v1 := e.Group("/v1")
	groupACustomer(v1)
	return e
}

// CUSTOMERS
func groupACustomer(e *echo.Group) {
	grA := e.Group("/customer")
	grA.GET("/set", controllers.HandlerSetSession)
	grA.GET("/all", controllers.HandlerGetCustomer)
	grA.GET("/delete", controllers.HandlerDeleteSession)
}
