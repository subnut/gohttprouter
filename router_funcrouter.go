package gohttprouter

import "net/http"

type funcRouter struct {
	router
}

func (r *router) FuncRouter() *funcRouter {
	og := r.Router()
	ret := funcRouter{*og}
	ret.parent = og
	return &ret
}

func (r *funcRouter) Route(m, p string, f func(http.ResponseWriter, *http.Request)) {
	r.routeAdd(m, p, http.HandlerFunc(f))
}

func (r *funcRouter) GET(p string, f func(http.ResponseWriter, *http.Request))     { r.Route("GET", p, f) }
func (r *funcRouter) HEAD(p string, f func(http.ResponseWriter, *http.Request))    { r.Route("HEAD", p, f) }
func (r *funcRouter) POST(p string, f func(http.ResponseWriter, *http.Request))    { r.Route("POST", p, f) }
func (r *funcRouter) PUT(p string, f func(http.ResponseWriter, *http.Request))     { r.Route("PUT", p, f) }
func (r *funcRouter) DELETE(p string, f func(http.ResponseWriter, *http.Request))  { r.Route("DELETE", p, f) }
func (r *funcRouter) OPTIONS(p string, f func(http.ResponseWriter, *http.Request)) { r.Route("OPTIONS", p, f) }
func (r *funcRouter) PATCH(p string, f func(http.ResponseWriter, *http.Request))   { r.Route("PATCH", p, f) }
