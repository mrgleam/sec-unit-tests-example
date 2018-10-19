package server

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mrgleam/sec-unit-tests-example/handlers"
)

func repository(s string) {
	fmt.Println(s)
}

func decorator(s string, e echo.HandlerFunc, t string) (string, echo.HandlerFunc) {
	repository(t)
	return s, e
}

func SetSecureMiddleWare() middleware.SecureConfig {
	return middleware.SecureConfig{
		XFrameOptions:      "DENY",
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
	}
}

func SetJWTMiddleWare() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("secret"),
		TokenLookup: "cookie:token",
	})
}

func EchoEngine(db *sql.DB) *echo.Echo {
	e := echo.New()
	// e.Use(middleware.SecureWithConfig(SetSecureMiddleWare()))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST(decorator("/logintest", handlers.Login(db), "sectesting.handlers.LoginRequestor"))
	e.File("/login.html", "public/login.html")
	e.File("/index.html", "public/index.html")

	e.POST("/api/login", handlers.Login(db))
	e.GET("/api/tasks", handlers.GetTasks(db), SetJWTMiddleWare())
	e.PUT("/api/tasks", handlers.PutTask(db), SetJWTMiddleWare())
	e.DELETE("/api/tasks/:id", handlers.DeleteTask(db), SetJWTMiddleWare())

	return e
}

func Server(db *sql.DB) {
	e := EchoEngine(db)
	e.Logger.Fatal(e.Start(":1323"))
}
