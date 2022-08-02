package gohttprouter

import "net/http"

type InterfaceRouter struct {
	r *Router
}

func (r *Router) FuncRouter() *FuncRouter { return &FuncRouter{r} }

func (r *InterfaceRouter) Router() *Router         { return r.r }
func (r *InterfaceRouter) FuncRouter() *FuncRouter { return r.r.FuncRouter() }

func (r *InterfaceRouter) Route(m string, p string, h http.Handler) {
	r.r.Route(m, p, h)
}
func (r *InterfaceRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.r.ServeHTTP(w, req)
}

func (r *InterfaceRouter) GET(p string, h http.Handler)     { r.Route("GET", p, h) }
func (r *InterfaceRouter) HEAD(p string, h http.Handler)    { r.Route("HEAD", p, h) }
func (r *InterfaceRouter) POST(p string, h http.Handler)    { r.Route("POST", p, h) }
func (r *InterfaceRouter) PUT(p string, h http.Handler)     { r.Route("PUT", p, h) }
func (r *InterfaceRouter) DELETE(p string, h http.Handler)  { r.Route("DELETE", p, h) }
func (r *InterfaceRouter) OPTIONS(p string, h http.Handler) { r.Route("OPTIONS", p, h) }
func (r *InterfaceRouter) PATCH(p string, h http.Handler)   { r.Route("PATCH", p, h) }
