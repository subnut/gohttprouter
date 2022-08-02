package gohttprouter

import "net/http"

type interfaceRouter struct {
	r *router
}

func (r *router) InterfaceRouter() *interfaceRouter { return &interfaceRouter{r} }

func (r *interfaceRouter) Router() *router           { return r.r }
func (r *interfaceRouter) FuncRouter() *funcRouter   { return r.r.FuncRouter() }

func (r *interfaceRouter) Handle(p string, h http.Handler) { r.r.Handle(p, h) }
func (r *interfaceRouter) Handler(req *http.Request) (handler http.Handler, pattern string) { return r.r.Handler(req) }
func (r *interfaceRouter) HandleFunc(p string, f func(http.ResponseWriter, *http.Request)) { r.r.HandleFunc(p, f) }
func (r *interfaceRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) { r.r.ServeHTTP(w, req) }

func (r *interfaceRouter) Route(m string, p string, h http.Handler) {
	r.r.routeAdd(m, p, h)
}

func (r *interfaceRouter) GET(p string, h http.Handler)     { r.Route("GET", p, h) }
func (r *interfaceRouter) HEAD(p string, h http.Handler)    { r.Route("HEAD", p, h) }
func (r *interfaceRouter) POST(p string, h http.Handler)    { r.Route("POST", p, h) }
func (r *interfaceRouter) PUT(p string, h http.Handler)     { r.Route("PUT", p, h) }
func (r *interfaceRouter) DELETE(p string, h http.Handler)  { r.Route("DELETE", p, h) }
func (r *interfaceRouter) OPTIONS(p string, h http.Handler) { r.Route("OPTIONS", p, h) }
func (r *interfaceRouter) PATCH(p string, h http.Handler)   { r.Route("PATCH", p, h) }
