package server

import (
	"fmt"
	"database/sql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mrgleam/sec-unit-tests-example/handlers"
)

func repository(s string) {
	fmt.Println(s)
}

func decorator(s string, e echo.HandlerFunc, t string) (string, echo.HandlerFunc) {
	repository(t)
	return s , e
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

	e.POST(decorator("/logintest", handlers.Login(db), "sectesting.handlers.LoginRequestor"))
	e.File("/login.html", "public/login.html")
	e.POST("/login", handlers.Login(db))

	r := e.Group("/restricted")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
		TokenLookup: "cookie:token",
	}))

	r.File("/index.html", "public/index.html")
	r.GET("/tasks", handlers.GetTasks(db))
	r.PUT("/tasks", handlers.PutTask(db))
	r.DELETE("/tasks/:id", handlers.DeleteTask(db))

	return e
}

func Server(db *sql.DB) {
	e := EchoEngine(db)
	e.Logger.Fatal(e.Start(":1323"))
}
