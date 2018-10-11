package server

import (
	"database/sql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mrgleam/sec-unit-tests-example/handlers"
)


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

	e.File("/login.html", "public/login.html")
	e.POST("/login", handlers.Login)

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
