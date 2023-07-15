package route

import (
	"fmt"

	"github.com/krzkro4122/echogogorm/controller"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func define_endpoints(e *echo.Echo) {
	// Index
	e.Static("/assets", "./public/assets")
	e.GET("/", controller.Index)

	// Authentication
	e.POST("/login", controller.Login)
	// e.POST("/register", controller.Register)
}

func Serve(address string) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	define_endpoints(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s", address)))
}

