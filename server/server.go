package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	User struct {
		Name  string `json:"name" form:"name"`
		Email string `json:"email" form:"email"`
	}
)

func GetUser(c echo.Context) error {
	email := c.Param("email")
	return c.JSON(http.StatusOK, email)
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get All Users")
}

func SetMiddleWareSecure() middleware.SecureConfig {
	return middleware.SecureConfig{
		XFrameOptions: "DENY",
	}
}

func NewEchoFramework() *echo.Echo {
	e := echo.New()
	e.Use(middleware.SecureWithConfig(SetMiddleWareSecure()))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users/:email", GetUser)
	e.GET("/users", GetUsers)
	return e
}

func Server() {
	e := NewEchoFramework()
	e.Logger.Fatal(e.Start(":1323"))
}
