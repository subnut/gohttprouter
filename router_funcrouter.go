package gohttprouter

import "net/http"

type funcRouter struct {
	router
	basicRouter
}

func (r *routers) FuncRouter() *funcRouter {
	return &funcRouter{*r.r, basicRouter{*r.r}}
}

func (r *funcRouter) Add(m string, p string, f handlerFunc) {
	r.routeAdd(m, p, http.HandlerFunc(f))
}

func (r *funcRouter) GET(p string, f handlerFunc)     { r.Add("GET", p, f) }
func (r *funcRouter) HEAD(p string, f handlerFunc)    { r.Add("HEAD", p, f) }
func (r *funcRouter) POST(p string, f handlerFunc)    { r.Add("POST", p, f) }
func (r *funcRouter) PUT(p string, f handlerFunc)     { r.Add("PUT", p, f) }
func (r *funcRouter) DELETE(p string, f handlerFunc)  { r.Add("DELETE", p, f) }
func (r *funcRouter) OPTIONS(p string, f handlerFunc) { r.Add("OPTIONS", p, f) }
func (r *funcRouter) PATCH(p string, f handlerFunc)   { r.Add("PATCH", p, f) }
