package gohttprouter

import (
	"fmt"
	"net/http"
)

type anyRouter struct {
	router
	basicRouter
}

func (r *routers) AnyRouter() *anyRouter {
	return &anyRouter{*r.r, basicRouter{*r.r}}
}

func (r *anyRouter) Add(method string, path string, handler any) {
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
func (r *anyRouter) GET(p string, f any)     { r.Add("GET", p, f) }
func (r *anyRouter) HEAD(p string, f any)    { r.Add("HEAD", p, f) }
func (r *anyRouter) POST(p string, f any)    { r.Add("POST", p, f) }
func (r *anyRouter) PUT(p string, f any)     { r.Add("PUT", p, f) }
func (r *anyRouter) DELETE(p string, f any)  { r.Add("DELETE", p, f) }
func (r *anyRouter) OPTIONS(p string, f any) { r.Add("OPTIONS", p, f) }
func (r *anyRouter) PATCH(p string, f any)   { r.Add("PATCH", p, f) }
