package gohttprouter

import "net/http"

type methodHandler struct {
	r *router
}

func (r *methodHandler) Route(m, p string, h http.Handler) {
	r.r.routeAdd(m, p, h)
}

func (r *methodHandler) GET(p string, h http.Handler)     { r.Route("GET", p, h) }
func (r *methodHandler) HEAD(p string, h http.Handler)    { r.Route("HEAD", p, h) }
func (r *methodHandler) POST(p string, h http.Handler)    { r.Route("POST", p, h) }
func (r *methodHandler) PUT(p string, h http.Handler)     { r.Route("PUT", p, h) }
func (r *methodHandler) DELETE(p string, h http.Handler)  { r.Route("DELETE", p, h) }
func (r *methodHandler) OPTIONS(p string, h http.Handler) { r.Route("OPTIONS", p, h) }
func (r *methodHandler) PATCH(p string, h http.Handler)   { r.Route("PATCH", p, h) }
