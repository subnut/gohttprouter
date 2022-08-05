package gohttprouter

import (
	"fmt"
	"net/http"
)

type runtimeRouter struct {
	router
	basicRouter
}

func (r *routers) RuntimeRouter() *runtimeRouter {
	return &runtimeRouter{*r.router, basicRouter{*r.router}}
}

func (r *runtimeRouter) Add(method string, path string, handler any) {
	if handler == nil {
		panic("nil handler")
	}
	switch handler.(type) {
	case http.Handler:
	case handlerFunc:
		handler = http.HandlerFunc(handler.(handlerFunc))
	default:
		panic(fmt.Sprintf("handler is of incompatible type %T\n"+
			"handler should be of type func(http.ResponseWriter, *http.Request) or http.Handler",
			handler))
	}
	r.routeAdd(method, path, handler.(http.Handler)) // <-- Type assertion!
}
func (r *runtimeRouter) GET(p string, f any)     { r.Add("GET", p, f) }
func (r *runtimeRouter) HEAD(p string, f any)    { r.Add("HEAD", p, f) }
func (r *runtimeRouter) POST(p string, f any)    { r.Add("POST", p, f) }
func (r *runtimeRouter) PUT(p string, f any)     { r.Add("PUT", p, f) }
func (r *runtimeRouter) DELETE(p string, f any)  { r.Add("DELETE", p, f) }
func (r *runtimeRouter) OPTIONS(p string, f any) { r.Add("OPTIONS", p, f) }
func (r *runtimeRouter) PATCH(p string, f any)   { r.Add("PATCH", p, f) }
