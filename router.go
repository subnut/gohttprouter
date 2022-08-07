package gohttprouter

import (
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	Config      config
	methods     []string
	middlewares []func(handlerFunc) handlerFunc
	routes      map[string]map[string]http.Handler
}
type config struct {
	Config
	Handlers *handlers
}
type Config struct {
	RoutePathAutoEncode   bool
	RoutePathIgnorePctEnc bool // `true` in http.ServeMux
	CaseInsensitive       bool
	KeepEmptySegments     bool
	KeepTrailingSlashes   bool // TODO
	LinkHeaderSkip        bool
	LinkHeaderHTTP        bool
	RedirectToCanonical   bool // TODO
	GlobalOPTIONShandler  handlerFunc
	DefaultOPTIONShandler handlerFunc
}

func (r *router) routeAdd(method string, path string, handler http.Handler) {
	if handler == nil {
		panic("nil handler")
	}
	path = r.normalizeRoutePath(path)
	r.routes[path] = make(map[string]http.Handler)
	r.routes[path][method] = handler
	for k, v := range r.methods {
		if v == method {
			break
		}
		if k+1 == len(r.methods) {
			// Whole array iterated, method not found.
			r.methods = append(r.methods, method)
		}
	}
}

func (r *router) getHandlers(request *http.Request) (map[string]http.Handler, string) {
	path := r.getPath(request)
	if r.Config.CaseInsensitive {
		path = strings.ToLower(path)
	}
	a, ok := r.routes[path]
	if ok != true {
		return nil, ""
	}
	return a, path
}
func (r *router) getHandler(request *http.Request) (handler http.Handler, pattern string) {
	handlers, path := r.getHandlers(request)
	if path != "" {
		if handler, exists := handlers[""]; exists == true {
			// Catch-all.
			return handler, path
		}
		if handler, exists := handlers[request.Method]; exists == true {
			return handler, path
		}
	}
	return http.HandlerFunc(r.Config.Handlers.NotFound), ""
}

func (r *router) serve(writer http.ResponseWriter, req *http.Request) {
	if req.Method == "OPTIONS" && req.RequestURI == "*" {
		return
	}
	h, path := r.getHandler(req)
	fun := h.ServeHTTP
	for _, middleware := range r.middlewares {
		fun = middleware(fun)
	}
	if !r.Config.LinkHeaderSkip && path != "" {
		scheme := "https"
		if r.Config.LinkHeaderHTTP {
			scheme = "http"
		}
		writer.Header().Add("Link", fmt.Sprintf(`<%s://%s%s>; rel="canonical"`, scheme, req.Host, path))
	}
	fun(writer, req)
}

// Validates the `path` argument string in router.routeAdd()
func (r *router) normalizeRoutePath(path string) string {
	if !r.Config.RoutePathIgnorePctEnc {
		// Validate percent-encoded characters in `path`
		hex := "0123456789" + "ABCDEF" + "abcdef"
		var invalidPctEncodings []string
		for i := 0; i < len(path); i++ {
			if path[i] != '%' {
				continue
			}
			if !(i+2 < len(path)) {
				invalidPctEncodings = append(invalidPctEncodings, path[i:])
				break
			}
			ishex := (strings.IndexByte(hex, path[i+1]) != -1) &&
				(strings.IndexByte(hex, path[i+2]) != -1)
			if !ishex {
				invalidPctEncodings = append(invalidPctEncodings, path[i:i+3])
			}
		}
		if len(invalidPctEncodings) != 0 {
			errmsg := "path contains invalid percent-encoded "
			if len(invalidPctEncodings) == 1 {
				errmsg += fmt.Sprintf("character %q", invalidPctEncodings[0])
			} else {
				errmsg += fmt.Sprintf("characters %q", invalidPctEncodings[0])
				for _, str := range invalidPctEncodings[1:] {
					errmsg += fmt.Sprintf(", %q", str)
				}
			}
			errmsg += fmt.Sprintf("\npath = %q", path)
			panic(errmsg)
		}
	}
	invalidChars := strings.Map(func(r rune) rune {
		if strings.IndexRune(rfc3986_pchar+"/", r) != -1 {
			return -1
		}
		return r
	}, path)
	if len(invalidChars) != 0 {
		if !r.Config.RoutePathAutoEncode {
			errmsg := fmt.Sprintf("path contains invalid characters %q", invalidChars)
			errmsg += fmt.Sprintf("\npath = %q", path)
			panic(errmsg)
		}
		path = encode(path)
	}
	if r.Config.CaseInsensitive {
		path = strings.ToLower(path)
	}
	return path
}
