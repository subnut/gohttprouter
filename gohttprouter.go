package gohttprouter

import (
	"net/http"
	"strings"
)

type routers struct {
	router *router
}

func Config() config {
	cfg := config{}
	cfg.RoutePathAutoEncode = true
	return cfg
}

func New() *routers {
	r := new(router)
	r.methods = []string{"OPTIONS"}
	r.routes = make(map[string]map[string]http.Handler)
	r.Responses = initResponses()
	r.config = Config()
	r.config.DefaultHandlers.GlobalOPTIONS =
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Allow", strings.Join(r.methods, ", "))
		}
	r.config.DefaultHandlers.DefaultOPTIONS =
		func(w http.ResponseWriter, req *http.Request) {
			a, p := r.getHandlers(req)
			if p == "" {
				r.Responses.NotFound(w, req)
				return
			}
			var methods []string
			for method := range a {
				methods = append(methods, method)
			}
			w.Header().Add("Allow", strings.Join(methods, ", "))
		}
	return &routers{r}
}

func (r *routers) WithConfig (cfg config) *routers {
	router := *r.router
	router.config = cfg
	return &routers{&router}
}
