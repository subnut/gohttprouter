package gohttprouter

import "net/http"

type funcRouter struct {
	r *router
}

func (r *router) FuncRouter() *funcRouter { return &funcRouter{r} }

func (r *funcRouter) Router() *router                   { return r.r }
func (r *funcRouter) InterfaceRouter() *interfaceRouter { return r.r.InterfaceRouter() }

func (r *funcRouter) Handle(p string, h http.Handler) { r.r.Handle(p, h) }
func (r *funcRouter) Handler(req *http.Request) (handler http.Handler, pattern string) { return r.r.Handler(req) }
func (r *funcRouter) HandleFunc(p string, f func(http.ResponseWriter, *http.Request)) { r.r.HandleFunc(p, f) }
func (r *funcRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) { r.r.ServeHTTP(w, req) }

func (r *funcRouter) Route(m string, p string, f func(http.ResponseWriter, *http.Request)) {
	r.r.routeAdd(m, p, http.HandlerFunc(f))
}

func (r *funcRouter) GET(p string, f func(http.ResponseWriter, *http.Request))     { r.Route("GET", p, f) }
func (r *funcRouter) HEAD(p string, f func(http.ResponseWriter, *http.Request))    { r.Route("HEAD", p, f) }
func (r *funcRouter) POST(p string, f func(http.ResponseWriter, *http.Request))    { r.Route("POST", p, f) }
func (r *funcRouter) PUT(p string, f func(http.ResponseWriter, *http.Request))     { r.Route("PUT", p, f) }
func (r *funcRouter) DELETE(p string, f func(http.ResponseWriter, *http.Request))  { r.Route("DELETE", p, f) }
func (r *funcRouter) OPTIONS(p string, f func(http.ResponseWriter, *http.Request)) { r.Route("OPTIONS", p, f) }
func (r *funcRouter) PATCH(p string, f func(http.ResponseWriter, *http.Request))   { r.Route("PATCH", p, f) }
