package server

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mrgleam/sec-unit-tests-example/handlers"
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

func SetSecureMiddleWare() middleware.SecureConfig {
	return middleware.SecureConfig{
		XFrameOptions:      "DENY",
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
	}
}

func EchoEngine(db *sql.DB) *echo.Echo {
	e := echo.New()
	e.Use(middleware.SecureWithConfig(SetSecureMiddleWare()))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/index.html", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))
	return e
}

func Server(db *sql.DB) {
	e := EchoEngine(db)
	e.Logger.Fatal(e.Start(":1323"))
}
