package gohttprouter

import "net/http"

type FuncRouter struct {
	r *Router
}

func (r *Router) InterfaceRouter() *InterfaceRouter { return &InterfaceRouter{r} }

func (r *FuncRouter) Router() *Router                   { return r.r }
func (r *FuncRouter) InterfaceRouter() *InterfaceRouter { return r.r.InterfaceRouter() }

func (r *FuncRouter) Route(m string, p string, f func(http.ResponseWriter, *http.Request)) {
	r.r.Route(m, p, http.HandlerFunc(f))
}
func (r *FuncRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.r.ServeHTTP(w, req)
}

func (r *FuncRouter) GET(p string, f func(http.ResponseWriter, *http.Request))     { r.Route("GET", p, f) }
func (r *FuncRouter) HEAD(p string, f func(http.ResponseWriter, *http.Request))    { r.Route("HEAD", p, f) }
func (r *FuncRouter) POST(p string, f func(http.ResponseWriter, *http.Request))    { r.Route("POST", p, f) }
func (r *FuncRouter) PUT(p string, f func(http.ResponseWriter, *http.Request))     { r.Route("PUT", p, f) }
func (r *FuncRouter) DELETE(p string, f func(http.ResponseWriter, *http.Request))  { r.Route("DELETE", p, f) }
func (r *FuncRouter) OPTIONS(p string, f func(http.ResponseWriter, *http.Request)) { r.Route("OPTIONS", p, f) }
func (r *FuncRouter) PATCH(p string, f func(http.ResponseWriter, *http.Request))   { r.Route("PATCH", p, f) }
