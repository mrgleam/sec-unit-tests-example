package method_support_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/mrgleam/sec-unit-tests-example/database"
	"github.com/mrgleam/sec-unit-tests-example/server"
	sec "github.com/mrgleam/sec-unit-tests-example/tests"
)

var AllMethod = []string{"GET", "PUT", "POST", "DELETE", "PATCH", "HEAD"}

var db = database.SetupDB()

func assertAllowMethod(t *testing.T, r gofight.HTTPResponse) {
	equals(t, 405, r.Code)
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func TestAllowMethod(t *testing.T) {
	e := server.EchoEngine(db)
	r := gofight.New()

	for _, v := range AllMethod {
		for k, c := range sec.RoutesChecker {
			if !contains(c.Method, v) {
				if v == "GET" {
					r.GET(k).
						Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
							assertAllowMethod(t, r)
						})
				} else if v == "PUT" {
					r.PUT(k).
						Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
							assertAllowMethod(t, r)
						})
				} else if v == "POST" {
					r.POST(k).
						Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
							assertAllowMethod(t, r)
						})
				} else if v == "DELETE" {
					r.DELETE(k).
						Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
							assertAllowMethod(t, r)
						})
				} else if v == "PATCH" {
					r.PATCH(k).
						Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
							assertAllowMethod(t, r)
						})
				} else if v == "HEAD" {
					r.HEAD(k).
						Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
							assertAllowMethod(t, r)
						})
				}
			}
		}
	}
	deleteDatabase()
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

func deleteDatabase() {
	path := "storage.db"
	err := os.Remove(path)

	if err != nil {
		fmt.Println(err)
	}
}
