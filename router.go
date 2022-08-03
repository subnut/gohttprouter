package gohttprouter

import (
	"fmt"
	"net/http"
)

// For compatibility with http.ServeMux
func (r *router) Handle(p string, h http.Handler) { r.routeAdd("", p, h) }
func (r *router) HandleFunc(p string, f handlerFunc) {
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
	F *funcHandler
	M *methodHandler
	config struct {
		EmptySegmentsAreImportant   bool
		TrailingSlashesAreImportant bool
	}
	middlewares []func(handlerFunc) handlerFunc
	routes      map[string]map[string]http.Handler
}

func New() *router {
	// Useful for setting default values that aren't the "nil" values
	r := new(router)
	r.F = &funcHandler{r}
	r.M = &methodHandler{r}
	r.routes = make(map[string]map[string]http.Handler)
	return r
}

func (r *router) Route(method string, path string, handler any) {
	if handler == nil {
		panic("nil handler")
	}
	switch handler.(type) {
	case http.Handler:
	case handlerFunc:
		handler = http.HandlerFunc(handler.(handlerFunc))
	default:
		panic(fmt.Sprintf("handler is of incompatible type %T\n"+
			"handler should be of type func(http.ResponseWriter, *http.Request) or http.Handler",
			handler))
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
