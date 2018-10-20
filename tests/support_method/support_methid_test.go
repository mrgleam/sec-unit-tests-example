package support_method_test

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

var AllMethod = []string{"GET", "PUT", "POST", "DELETE", "TRACE", "HEAD"}

var db = database.SetupDB()

func assertAllowMethod(t *testing.T, r gofight.HTTPResponse) {
	equals(t, 405, r.Code)
}

func TestAllowMethod(t *testing.T) {
	e := server.EchoEngine(db)
	r := gofight.New()

	for k, c := range sec.RoutesChecker {
		for _, m := range c.Method {
			if m == "GET" {
				r.GET(k).
					Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
						assertAllowMethod(t, r)
					})
			} else if m == "PUT" {
				r.PUT(k).
					Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
						assertAllowMethod(t, r)
					})
			} else if m == "POST" {
				r.POST(k).
					Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
						assertAllowMethod(t, r)
					})
			} else if m == "DELETE" {
				r.DELETE(k).
					Run(e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
						assertAllowMethod(t, r)
					})
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
