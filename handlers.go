package gohttprouter

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type responses struct {
	codes []int
	numResponses
	namedResponses
}
type numResponses struct {
	SC201 handlerFunc // Created
	SC307 handlerFunc // Temporary Redirect
	SC308 handlerFunc // Permanent Redirect
	SC404 handlerFunc // Not Found
	SC405 handlerFunc // Method Not Allowed
}
type namedResponses struct {
	Created           handlerFunc `sc:"201"`
	TemporaryRedirect handlerFunc `sc:"307"`
	PermanentRedirect handlerFunc `sc:"308"`
	NotFound          handlerFunc `sc:"404"`
	MethodNotAllowed  handlerFunc `sc:"405"`
}

func (h *responses) Set(sc int, handler handlerFunc) {
	found := false
	for _, code := range h.codes {
		if sc == code {
			found = true
			break
		}
	}
	if !found {
		panic(fmt.Sprintf("invalid Status Code %v", sc))
	}
	s := reflect.ValueOf(h).Elem()
	s.FieldByName(fmt.Sprintf("SC%d", sc)).Set(reflect.ValueOf(handler))
	t := reflect.TypeOf(namedResponses{})
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("sc") == strconv.Itoa(sc) {
			s.FieldByName(t.Field(i).Name).Set(reflect.ValueOf(handler))
			break
		}
	}
}

func initResponses() *responses {
	var codes []int
	t := reflect.TypeOf(numResponses{})
	for i := 0; i < t.NumField(); i++ {
		code, _ := strconv.Atoi(strings.TrimLeft(t.Field(i).Name, "SC"))
		codes = append(codes, code)
	}
	num := numResponses{}
	nam := namedResponses{}
	num.init()
	nam.init()
	return &responses{codes, num, nam}
}
func (h *numResponses) init() {
	t := reflect.TypeOf(*h)
	s := reflect.ValueOf(h).Elem()
	for i := 0; i < s.NumField(); i++ {
		s.Field(i).Set(reflect.ValueOf(func(SC int, e error) handlerFunc {
			return func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(SC) }
		}(strconv.Atoi(strings.TrimLeft(t.Field(i).Name, "SC")))))
	}
}
func (h *namedResponses) init() {
	t := reflect.TypeOf(*h)
	s := reflect.ValueOf(h).Elem()
	for i := 0; i < s.NumField(); i++ {
		s.Field(i).Set(reflect.ValueOf(func(SC int, e error) handlerFunc {
			return func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(SC) }
		}(strconv.Atoi(t.Field(i).Tag.Get("sc")))))
	}
}
