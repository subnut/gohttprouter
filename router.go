package gohttprouter

import "net/http"

// For compatibility with http.ServeMux
func (r *router) Handle(p string, h http.Handler) { r.routeAdd("", p, h) }
func (r *router) HandleFunc(p string, f func(http.ResponseWriter, *http.Request)) {
	r.routeAdd("", p, http.HandlerFunc(f))
}
func (r *router) Handler(req *http.Request) (handler http.Handler, pattern string) {
	// TODO: Think about it.
	return nil, ""
}
func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.serve(w, req)
}

type router struct {
	config struct {
		EmptySegmentsAreImportant   bool
		TrailingSlashesAreImportant bool
	}
	middlewares []func(http.ResponseWriter, *http.Request) func(http.ResponseWriter, *http.Request)
	routes      map[string]map[string]http.Handler
}

func (r *router) Route(method string, path string, handler any) {
	if handler == nil {
		panic("nil handler")
	}
	switch handler.(type) {
	case http.Handler:
	case func(http.ResponseWriter, *http.Request):
		handler = http.HandlerFunc(handler.(func(http.ResponseWriter, *http.Request)))
	default:
		panic("handler is of incompatible type")
	}
	r.routeAdd(method, path, handler.(http.Handler)) // <-- Type assertion!
}
func (r *router) routeAdd(method string, path string, handler http.Handler) {
	if handler == nil {
		panic("nil handler")
	}
	r.routes[path] = make(map[string]http.Handler)
	r.routes[path][method] = handler
}

func (r *router) serve(writer http.ResponseWriter, request *http.Request) {
	a, ok := r.routes[r.getPath(request)]
	if ok != true {
		http.NotFound(writer, request)
		return
	}
	if _, exists := a[""]; exists == true {
		// Catch-all.
		a[""].ServeHTTP(writer, request)
	} else {
		a[request.Method].ServeHTTP(writer, request)
	}
}
