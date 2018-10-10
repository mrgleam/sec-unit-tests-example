package server

import (
	"fmt"
	"testing"

	"github.com/labstack/echo"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func getRoutes() []*echo.Route {
	e := NewEchoFramework()
	return e.Routes()
}
func TestWebSecureHeaderInclusionPolicy(t *testing.T) {
	e := getRoutes()
	for _, route := range e {
		fmt.Println("Path:", route.Path)
		fmt.Println("Method:", route.Method)
		r := gofight.New()

		if route.Method == "GET" {
			r.GET(route.Path).
				Run(NewEchoFramework(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
					assert.Equal(t, []string{"DENY"}, r.HeaderMap["X-Frame-Options"])
					assert.Equal(t, []string{"nosniff"}, r.HeaderMap["X-Content-Type-Options"])
					assert.Equal(t, []string{"1; mode=block"}, r.HeaderMap["X-Xss-Protection"])
				})
		}
	}
}
