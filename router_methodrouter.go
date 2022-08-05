package gohttprouter

import "net/http"

type methodRouter struct {
	router
	basicRouter
}

func (r *routers) MethodRouter() *methodRouter {
	return &methodRouter{*r.router, basicRouter{*r.router}}
}

func (r *methodRouter) Add(m string, p string, f http.Handler) {
	r.routeAdd(m, p, f)
}

func (r *methodRouter) GET(p string, f http.Handler)     { r.Add("GET", p, f) }
func (r *methodRouter) HEAD(p string, f http.Handler)    { r.Add("HEAD", p, f) }
func (r *methodRouter) POST(p string, f http.Handler)    { r.Add("POST", p, f) }
func (r *methodRouter) PUT(p string, f http.Handler)     { r.Add("PUT", p, f) }
func (r *methodRouter) DELETE(p string, f http.Handler)  { r.Add("DELETE", p, f) }
func (r *methodRouter) OPTIONS(p string, f http.Handler) { r.Add("OPTIONS", p, f) }
func (r *methodRouter) PATCH(p string, f http.Handler)   { r.Add("PATCH", p, f) }
