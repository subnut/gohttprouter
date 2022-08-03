package gohttprouter

import "net/http"

type basicRouter struct {
	router
}

func (r *routers) BasicRouter() *basicRouter {
	return &basicRouter{*r.r}
}

func (r *basicRouter) Handle(p string, h http.Handler) {
	r.routeAdd("", p, h)
}
func (r *basicRouter) HandleFunc(p string, f handlerFunc) {
	r.routeAdd("", p, http.HandlerFunc(f))
}
func (r *basicRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.serve(w, req)
}
func (r *basicRouter) Handler(req *http.Request) (handler http.Handler, pattern string) {
	// TODO: Think about it.
	return nil, ""
}