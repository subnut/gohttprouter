package gohttprouter

import "net/http"

func New() *Router {
	// Useful for setting default values that aren't the "nil" values
	router := new(Router)
	router.routes = make(map[string]map[string]http.Handler)
	return router
}
