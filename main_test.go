package main_test

import (
	"fmt"
	"testing"

	"github.com/labstack/echo"
	"github.com/mrgleam/sec-unit-tests-example/database"
	"github.com/mrgleam/sec-unit-tests-example/server"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

var db = database.SetupDB()

func getRoutes() []*echo.Route {
	e := server.EchoEngine(db)
	return e.Routes()
}
func assertHeaderInclusionPolicy(t *testing.T, r gofight.HTTPResponse) {
	assert.Equal(t, []string{"DENY"}, r.HeaderMap["X-Frame-Options"])
	assert.Equal(t, []string{"nosniff"}, r.HeaderMap["X-Content-Type-Options"])
	assert.Equal(t, []string{"1; mode=block"}, r.HeaderMap["X-Xss-Protection"])
}
func TestWebSecureHeaderInclusionPolicy(t *testing.T) {
	e := getRoutes()
	for _, route := range e {
		fmt.Println("Path:", route.Path)
		fmt.Println("Method:", route.Method)
		r := gofight.New()

		if route.Method == "GET" {
			r.GET(route.Path).
				Run(server.EchoEngine(db), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
					assertHeaderInclusionPolicy(t, r)
				})
		} else if route.Method == "DELETE" {
			r.DELETE(route.Path).
				Run(server.EchoEngine(db), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
					assertHeaderInclusionPolicy(t, r)
				})
		} else if route.Method == "PUT" {
			r.PUT(route.Path).
				Run(server.EchoEngine(db), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
					assertHeaderInclusionPolicy(t, r)
				})
		} else if route.Method == "POST" {
			r.PUT(route.Path).
				Run(server.EchoEngine(db), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
					assertHeaderInclusionPolicy(t, r)
				})
		}
	}
}
