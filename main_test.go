package main_test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/buger/jsonparser"
	"github.com/labstack/echo"

	"github.com/mrgleam/sec-unit-tests-example/database"
	"github.com/mrgleam/sec-unit-tests-example/server"
	sec "github.com/mrgleam/sec-unit-tests-example/tests"

	"github.com/appleboy/gofight"
)

var db = database.SetupDB()

func assertAuthentication(t *testing.T, r gofight.HTTPResponse) {
	equals(t, 400, r.Code)
}

func assertHeaderInclusionPolicy(t *testing.T, r gofight.HTTPResponse) {
	equals(t, []string{"DENY"}, r.HeaderMap["X-Frame-Options"])
	equals(t, []string{"nosniff"}, r.HeaderMap["X-Content-Type-Options"])
	equals(t, []string{"1; mode=block"}, r.HeaderMap["X-Xss-Protection"])
}
func TestWebSecureHeaderInclusionPolicy(t *testing.T) {
	e := server.EchoEngine(db)
	r := gofight.New()
	token := LoginWitTestData(r, e)

	for _, route := range e.Routes() {
		fmt.Println("Path:", route.Path)
		fmt.Println("Method:", route.Method)

		if sec.RoutesChecker[route.Path].RequireAuthen {
			if route.Method == "GET" {
				r.GET(route.Path).
					SetCookie(gofight.H{
						"token": token,
					}).
					Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
						assertHeaderInclusionPolicy(t, r)
					})
			} else if route.Method == "DELETE" {
				r.DELETE(route.Path).
					SetCookie(gofight.H{
						"token": token,
					}).
					Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
						assertHeaderInclusionPolicy(t, r)
					})
			} else if route.Method == "PUT" {
				r.PUT(route.Path).
					SetCookie(gofight.H{
						"token": token,
					}).
					Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
						assertHeaderInclusionPolicy(t, r)
					})
			} else if route.Method == "POST" {
				r.POST(route.Path).
					SetCookie(gofight.H{
						"token": token,
					}).
					Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
						assertHeaderInclusionPolicy(t, r)
					})
			}
		} else {
			if route.Method == "GET" {
				r.GET(route.Path).
					Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
						assertHeaderInclusionPolicy(t, r)
					})
			} else if route.Method == "DELETE" {
				r.DELETE(route.Path).
					Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
						assertHeaderInclusionPolicy(t, r)
					})
			} else if route.Method == "PUT" {
				r.PUT(route.Path).
					Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
						assertHeaderInclusionPolicy(t, r)
					})
			} else if route.Method == "POST" {
				r.POST(route.Path).
					Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
						assertHeaderInclusionPolicy(t, r)
					})
			}
		}
	}
}

func TestAuthentication(t *testing.T) {
	e := server.EchoEngine(db)
	r := gofight.New()

	for k, c := range sec.RoutesChecker {
		if c.RequireAuthen {
			for _, m := range c.Method {
				fmt.Println("Method", m)
				if m == "GET" {
					r.GET(k).
						Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
							assertAuthentication(t, r)
						})
				} else if m == "PUT" {
					r.PUT(k).
						Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
							assertAuthentication(t, r)
						})
				} else if m == "POST" {
					r.POST(k).
						Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
							assertAuthentication(t, r)
						})
				} else if m == "DELETE" {
					r.DELETE(k).
						Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
							assertAuthentication(t, r)
						})
				}
			}
		}
	}
}

func LoginWitTestData(r *gofight.RequestConfig, e *echo.Echo) string {
	var token string
	r.POST("/api/login").
		SetJSON(gofight.D{
			"email":    "test01@test.com",
			"password": "test01",
		}).
		Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(r.Body.String())

			token, _ = jsonparser.GetString(data, "token")
		})
	return token
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
