package gohttprouter

import "net/http"

type interfaceRouter struct {
	router
}

func (r *router) InterfaceRouter() *interfaceRouter {
	og := r.Router()
	ret := interfaceRouter{*og}
	ret.parent = og
	return &ret
}

func (r *interfaceRouter) Route(m, p string, h http.Handler) {
	r.routeAdd(m, p, h)
}

func (r *interfaceRouter) GET(p string, h http.Handler)     { r.Route("GET", p, h) }
func (r *interfaceRouter) HEAD(p string, h http.Handler)    { r.Route("HEAD", p, h) }
func (r *interfaceRouter) POST(p string, h http.Handler)    { r.Route("POST", p, h) }
func (r *interfaceRouter) PUT(p string, h http.Handler)     { r.Route("PUT", p, h) }
func (r *interfaceRouter) DELETE(p string, h http.Handler)  { r.Route("DELETE", p, h) }
func (r *interfaceRouter) OPTIONS(p string, h http.Handler) { r.Route("OPTIONS", p, h) }
func (r *interfaceRouter) PATCH(p string, h http.Handler)   { r.Route("PATCH", p, h) }
