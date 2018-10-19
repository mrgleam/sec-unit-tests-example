package sec

import "github.com/mrgleam/sec-unit-tests-example/models"

type Checker struct {
	Method        []string
	RequireAuthen bool
	Models        interface{}
}

var RoutesChecker = map[string]Checker{
	"/login":      Checker{[]string{"POST"}, false, models.User{}},
	"/logintest":  Checker{[]string{"POST"}, false, models.User{}},
	"/login.html": Checker{[]string{"GET"}, false, nil},
	"/api/tasks":  Checker{[]string{"GET", "PUT", "DELETE"}, true, models.Task{}},
}
