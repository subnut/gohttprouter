package gohttprouter

import (
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	config      config
	methods     []string
	middlewares []func(handlerFunc) handlerFunc
	routes      map[string]map[string]http.Handler
	Responses   *responses
}
type config struct {
	RoutePathAutoEncode   bool
	RoutePathIgnorePctEnc bool // `true` in http.ServeMux
	CaseInsensitive       bool
	KeepEmptySegments     bool
	KeepTrailingSlashes   bool // TODO
	LinkHeaderSkip        bool
	LinkHeaderHTTP        bool
	RedirectToCanonical   bool
	DefaultHandlers       defaultHandlers
}
type defaultHandlers struct {
	DefaultOPTIONS handlerFunc
	GlobalOPTIONS  handlerFunc
	PanicHandler   func(any) handlerFunc // TODO
	// PanicHandler shall consume the panic, and shall return a handlerFunc
	// which (when called) shall preferably respond with an HTTP 500 response
	// containing details of what went wrong (or a brief error message)
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
	if r.config.CaseInsensitive {
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
	return nil, ""
}

func (r *router) serve(writer http.ResponseWriter, req *http.Request) {
	if req.Method == "OPTIONS" && req.RequestURI == "*" {
		r.config.DefaultHandlers.GlobalOPTIONS(writer, req)
		return
	}
	h, path := r.getHandler(req)
	if path == "" {
		switch req.Method {
		case "OPTIONS":
			r.config.DefaultHandlers.DefaultOPTIONS(writer, req)
			return
		case "HEAD":
			rq := http.Request(*req)
			rq.Method = "GET"
			h, path = r.getHandler(&rq)
			if path != "" {
				break
			}
			fallthrough
		default:
			r.Responses.NotFound(writer, req)
			return
		}
	}
	requestPath := req.RequestURI
	if i := strings.IndexByte(requestPath, '?'); i != -1 {
		requestPath = requestPath[:i]
	}
	if i := strings.IndexByte(requestPath, '#'); i != -1 {
		requestPath = requestPath[:i]
	}
	if path != requestPath && r.config.RedirectToCanonical {
		r.Responses.PermanentRedirect(writer, req)
		writer.Header().Add("Location", path)
		return
	}
	fun := h.ServeHTTP
	for _, middleware := range r.middlewares {
		fun = middleware(fun)
	}
	if !r.config.LinkHeaderSkip && path != "" {
		scheme := "https"
		if r.config.LinkHeaderHTTP {
			scheme = "http"
		}
		writer.Header().Add("Link", fmt.Sprintf(`<%s://%s%s>; rel="canonical"`, scheme, req.Host, path))
	}
	fun(writer, req)
}

// Validates the `path` argument string in router.routeAdd()
func (r *router) normalizeRoutePath(path string) string {
	if !r.config.RoutePathIgnorePctEnc {
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
		if !r.config.RoutePathAutoEncode {
			errmsg := fmt.Sprintf("path contains invalid characters %q", invalidChars)
			errmsg += fmt.Sprintf("\npath = %q", path)
			panic(errmsg)
		}
		path = encode(path)
	}
	if r.config.CaseInsensitive {
		path = strings.ToLower(path)
	}
	return path
}
