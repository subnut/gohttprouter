package gohttprouter

import (
	"net/http"
	"strings"
)

type routers struct {
	router *router
}

func New() *routers {
	r := new(router)
	r.methods = []string{"OPTIONS"}
	r.routes = make(map[string]map[string]http.Handler)
	r.Config.RoutePathAutoEncode = true
	r.Config.Handlers = initHandlers()
	r.Config.GlobalOPTIONShandler = func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Allow", strings.Join(r.methods, ", "))
	}
	r.Config.DefaultOPTIONShandler = func(w http.ResponseWriter, req *http.Request) {
		a, p := r.getHandlers(req)
		if p == "" {
			r.Config.Handlers.NotFound(w, req)
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
