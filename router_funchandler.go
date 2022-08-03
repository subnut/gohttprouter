package gohttprouter

import "net/http"

type funcHandler struct {
	r *router
}

func (r *funcHandler) Route(m string, p string, f handlerFunc) {
	r.r.routeAdd(m, p, http.HandlerFunc(f))
}

func (r *funcHandler) GET(p string, f handlerFunc)     { r.Route("GET", p, f) }
func (r *funcHandler) HEAD(p string, f handlerFunc)    { r.Route("HEAD", p, f) }
func (r *funcHandler) POST(p string, f handlerFunc)    { r.Route("POST", p, f) }
func (r *funcHandler) PUT(p string, f handlerFunc)     { r.Route("PUT", p, f) }
func (r *funcHandler) DELETE(p string, f handlerFunc)  { r.Route("DELETE", p, f) }
func (r *funcHandler) OPTIONS(p string, f handlerFunc) { r.Route("OPTIONS", p, f) }
func (r *funcHandler) PATCH(p string, f handlerFunc)   { r.Route("PATCH", p, f) }
