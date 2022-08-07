package gohttprouter

import "net/http"

type router struct {
	config struct {
		EmptySegmentsAreImportant   bool
		TrailingSlashesAreImportant bool
	}
	middlewares []func(handlerFunc) handlerFunc
	routes      map[string]map[string]http.Handler
}

func (r *router) routeAdd(method string, path string, handler http.Handler) {
	if handler == nil {
		panic("nil handler")
	}
	r.routes[path] = make(map[string]http.Handler)
	r.routes[path][method] = handler
}

func (r *router) getHandler(request *http.Request) (handler http.Handler, pattern string) {
	path := r.getPath(request)
	a, ok := r.routes[path]
	if ok != true {
		return http.NotFoundHandler(), ""
	}
	if _, exists := a[""]; exists == true {
		// Catch-all.
		return a[""], path
	}
	if _, exists := a[request.Method]; exists == true {
		return a[request.Method], path
	}
	return http.NotFoundHandler(), ""
}

func (r *router) serve(writer http.ResponseWriter, request *http.Request) {
	h, _ := r.getHandler(request)
	fun := h.ServeHTTP
	// Use middlewares
	for _, middleware := range r.middlewares {
		fun = middleware(fun)
	}
	fun(writer, request)
}
