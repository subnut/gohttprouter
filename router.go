package gohttprouter

import (
	"net/http"

	"github.com/subnut/gohttprouter/utils"
)

type Router struct {
	config config
}

/*
	Preserve is false by default
	Redirect is true by default

	Examples
	========
	- PreserveTrailingSlash: true
		/hi/ => call handler for /hi/, 404 if no handler
		/hi  => call handler for /hi,  404 if no handler

	- PreserveTrailingSlash: false
	  RedirectTrailingSlash: false
		/hi/ => call handler for /hi/,
		        else, call handler for /hi (if defined)
		        else, 404
		/hi =>  call handler for /hi,
		        else, call handler for /hi/ (if defined)
		        else, 404

	- PreserveTrailingSlash: false
	  RedirectTrailingSlash: true
		/hi/ => call handler for /hi/,
		        else, REDIRECT to /hi (if /hi handler is defined)
		        else, 404
		/hi =>  call handler for /hi,
		        else, REDIRECT to /hi/ (if /hi/ handler is defined)
		        else, 404
*/
type config struct {
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
