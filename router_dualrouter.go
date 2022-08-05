package gohttprouter

import "net/http"

type dualRouter struct {
	router
	basicRouter
	F *funcHandler
	I *interfaceHandler
}

func (r *routers) DualRouter() *dualRouter {
	return &dualRouter{
		*r.router,
		basicRouter{*r.router},
		&funcHandler{r.router},
		&interfaceHandler{r.router},
	}
}

type funcHandler struct{ r *router }
type interfaceHandler struct{ r *router }

func (r *funcHandler) Add(m string, p string, f handlerFunc) {
	r.r.routeAdd(m, p, http.HandlerFunc(f))
}
func (r *funcHandler) GET(p string, f handlerFunc)     { r.Add("GET", p, f) }
func (r *funcHandler) HEAD(p string, f handlerFunc)    { r.Add("HEAD", p, f) }
func (r *funcHandler) POST(p string, f handlerFunc)    { r.Add("POST", p, f) }
func (r *funcHandler) PUT(p string, f handlerFunc)     { r.Add("PUT", p, f) }
func (r *funcHandler) DELETE(p string, f handlerFunc)  { r.Add("DELETE", p, f) }
func (r *funcHandler) OPTIONS(p string, f handlerFunc) { r.Add("OPTIONS", p, f) }
func (r *funcHandler) PATCH(p string, f handlerFunc)   { r.Add("PATCH", p, f) }

func (r *interfaceHandler) Add(m, p string, h http.Handler) {
	r.r.routeAdd(m, p, h)
}
func (r *interfaceHandler) GET(p string, h http.Handler)     { r.Add("GET", p, h) }
func (r *interfaceHandler) HEAD(p string, h http.Handler)    { r.Add("HEAD", p, h) }
func (r *interfaceHandler) POST(p string, h http.Handler)    { r.Add("POST", p, h) }
func (r *interfaceHandler) PUT(p string, h http.Handler)     { r.Add("PUT", p, h) }
func (r *interfaceHandler) DELETE(p string, h http.Handler)  { r.Add("DELETE", p, h) }
func (r *interfaceHandler) OPTIONS(p string, h http.Handler) { r.Add("OPTIONS", p, h) }
func (r *interfaceHandler) PATCH(p string, h http.Handler)   { r.Add("PATCH", p, h) }
