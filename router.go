package gohttprouter

import (
	"net/http"
)

// Convenience functions
//
// 	RFC 9110 § 9	 GET, HEAD, POST, PUT, DELETE, OPTIONS
// 	RFC 5789	 PATCH
//
// See:
// 	https://www.iana.org/assignments/http-methods/http-methods.xhtml
// 	https://datatracker.ietf.org/doc/html/rfc9110
// 	https://datatracker.ietf.org/doc/html/rfc5789
func (r *Router) GET(p string, h http.HandlerFunc)     { r.Route("GET", p, h) }
func (r *Router) HEAD(p string, h http.HandlerFunc)    { r.Route("HEAD", p, h) }
func (r *Router) POST(p string, h http.HandlerFunc)    { r.Route("POST", p, h) }
func (r *Router) PUT(p string, h http.HandlerFunc)     { r.Route("PUT", p, h) }
func (r *Router) DELETE(p string, h http.HandlerFunc)  { r.Route("DELETE", p, h) }
func (r *Router) OPTIONS(p string, h http.HandlerFunc) { r.Route("OPTIONS", p, h) }
func (r *Router) PATCH(p string, h http.HandlerFunc)   { r.Route("PATCH", p, h) }

type Router struct {
	config struct {
		EmptySegmentsAreImportant   bool
		TrailingSlashesAreImportant bool
	}
	middlewares []func(http.HandlerFunc) http.HandlerFunc
}

func New() *Router {
	// Useful for setting default values that aren't the "nil" values
	return &Router{}
}

func (*Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
}

func (*Router) Route(method string, path string, handler http.HandlerFunc) {
}
