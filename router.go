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

func (r *router) serve(writer http.ResponseWriter, request *http.Request) {
	var fun handlerFunc
	a, ok := r.routes[r.getPath(request)]
	if ok != true {
		http.NotFound(writer, request)
		return
	}
	if _, exists := a[""]; exists == true {
		// Catch-all.
		fun = a[""].ServeHTTP
	} else {
		fun = a[request.Method].ServeHTTP
	}
	// Use middlewares
	for _, mitm := range r.middlewares {
		fun = mitm(fun)
	}
	fun(writer, request)
}
