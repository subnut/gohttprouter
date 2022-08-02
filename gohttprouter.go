package gohttprouter

import "net/http"

func New() *router {
	// Useful for setting default values that aren't the "nil" values
	r := new(router)
	r.routes = make(map[string]map[string]http.Handler)
	return r
}
