package gohttprouter

import "net/http"

// For compatiblity with http.ServeMux
func (r *Router) Handle(p string, h http.Handler) { r.Route("", p, h) }
func (r *Router) HandleFunc(p string, f func(http.ResponseWriter, *http.Request)) {
	r.Route("", p, http.HandlerFunc(f))
}

type Router struct {
	config struct {
		EmptySegmentsAreImportant   bool
		TrailingSlashesAreImportant bool
	}
	middlewares []func(http.ResponseWriter, *http.Request) func(http.ResponseWriter, *http.Request)
	routes      map[string]map[string]http.Handler
}

func (router *Router) Route(method string, path string, handler http.Handler) {
	if handler == nil {
		panic("nil handler")
	}
	router.routes[path] = make(map[string]http.Handler)
	router.routes[path][method] = handler
}

func (router *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	a, ok := router.routes[router.getPath(request)]
	if ok != true {
		http.NotFound(writer, request)
		return
	}
	if _,exists := a[""]; exists == true {
		// Catch-all.
		a[""].ServeHTTP(writer, request)
	} else {
		a[request.Method].ServeHTTP(writer, request)
	}
}
