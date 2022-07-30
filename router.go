package gohttprouter

import (
	"net/http"
)

type Router struct {
	config config
}

type config struct {
	/*
	   Preserve is false by default
	   Redirect is true by default

	   Example for TrailingSlash -
	   +----------+----------+------+----------------------------+
	   | Preserve | Redirect | URL  | Action                     |
	   +----------+----------+------+----------------------------+
	   | false    | false    | /hi/ | Call the handler for /hi   |
	   | false    | true     | /hi/ | Redirect to /hi            |
	   | true     | *        | /hi/ | 404 if no handler for /hi/ |
	   +----------+----------+------+----------------------------+
	*/
	PreserveEmptySegments bool
	RedirectEmptySegments bool
	PreserveTrailingSlash bool
	RedirectTrailingSlash bool
}

func New() *Router {
	return &Router{
		config: config{
			RedirectTrailingSlash: true,
			RedirectEmptySegments: true,
		},
	}
}

func (*Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
}
