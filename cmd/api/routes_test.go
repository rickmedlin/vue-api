package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_Routes_Exist(t *testing.T) {
	testRoutes := testApp.routes()
	chiRoutes := testRoutes.(chi.Router)

	// these routes must exist
	routeExists(t, chiRoutes, "/users/login")
	routeExists(t, chiRoutes, "/users/login")
	routeExists(t, chiRoutes, "/admin/users/get/{id}")
	routeExists(t, chiRoutes, "/admin/users/get/{id}")
	routeExists(t, chiRoutes, "/admin/users")
	routeExists(t, chiRoutes, "/admin/users/delete")
}

func routeExists(t *testing.T, routes chi.Router, route string) {
	// assume route does not exist
	found := false

	// walk through all routes
	_ = chi.Walk(routes, func(method string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		// if route is found, set to true
		if route == foundRoute {
			found = true
		}

		return nil
	})

	// return an error if route is not found
	if !found {
		t.Errorf("did not find %s in registered routes.", route)
	}

}
