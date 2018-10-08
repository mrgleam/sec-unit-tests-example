package server

import (
	"fmt"
	"testing"

	"github.com/labstack/echo"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func getRoutes() []*echo.Route {
	// Setup
	e := NewEchoFramework()
	// data, err := json.MarshalIndent(e.Routes(), "", "  ")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("data", string(data))

	// for _, route := range e.Routes() {
	// 	fmt.Println("Path:", route.Path)
	// 	fmt.Println("Method:", route.Method)
	// }
	return e.Routes()
}
func TestWebSecurityHeaderPolicy(t *testing.T) {
	e := getRoutes()
	for _, route := range e {
		fmt.Println("Path:", route.Path)
		fmt.Println("Method:", route.Method)
		r := gofight.New()

		if route.Method == "GET" {
			r.GET(route.Path).
			Run(NewEchoFramework(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, []string{"DENY"}, r.HeaderMap["X-Frame-Options"])
			})
		}
	}
}

// var (
// 	emailString = `"jon@labstack.com"`
// )

// func TestGetUser(t *testing.T) {
// 	// Setup
// 	e := NewEchoFramework()
// 	data, err := json.MarshalIndent(e.Routes(), "", "  ")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("data", string(data))
// 	req := httptest.NewRequest("GET", "/users/:email", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetParamNames("email")
// 	c.SetParamValues("jon@labstack.com")
// 	// middleware.SecureWithConfig(SetMiddleWareSecure())(GetUser)(c)
// 	// fmt.Println("request", c.Request())
// 	// Assertions
// 	// if assert.NoError(t, GetUser(c)) {
// 	// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	// 	assert.Equal(t, true, string(rec.Header))
// 	// 	assert.Equal(t, emailString, rec.Body.String())
// 	// }
// 	// GetUser(c)
// 	if ctype := rec.Header().Get("X-Frame-Options"); ctype != "DENY" {
// 		t.Errorf("content type header X-Frame-Options does not match: got %#v want %#v", ctype, "DENY")
// 	}
// }
