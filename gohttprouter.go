package gohttprouter

import "net/http"

type routers struct {
	r *router
}

func New() *routers {
	// Useful for setting default values that aren't the "nil" values
	r := new(router)
	r.routes = make(map[string]map[string]http.Handler)
	return &routers{r}
}
