package gohttprouter

import "net/http"

type interfaceRouter struct {
	router
	basicRouter
}

func (r *routers) InterfaceRouter() *interfaceRouter {
	return &interfaceRouter{*r.r, basicRouter{*r.r}}
}

func (r *interfaceRouter) Add(m string, p string, f http.Handler) {
	r.routeAdd(m, p, f)
}

func (r *interfaceRouter) GET(p string, f http.Handler)     { r.Add("GET", p, f) }
func (r *interfaceRouter) HEAD(p string, f http.Handler)    { r.Add("HEAD", p, f) }
func (r *interfaceRouter) POST(p string, f http.Handler)    { r.Add("POST", p, f) }
func (r *interfaceRouter) PUT(p string, f http.Handler)     { r.Add("PUT", p, f) }
func (r *interfaceRouter) DELETE(p string, f http.Handler)  { r.Add("DELETE", p, f) }
func (r *interfaceRouter) OPTIONS(p string, f http.Handler) { r.Add("OPTIONS", p, f) }
func (r *interfaceRouter) PATCH(p string, f http.Handler)   { r.Add("PATCH", p, f) }
