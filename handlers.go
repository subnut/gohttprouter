package gohttprouter

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type handlers struct {
	codes []int
	numHandlers
	namedHandlers
}
type numHandlers struct {
	SC201 handlerFunc // Created
	SC404 handlerFunc // Not Found
	SC405 handlerFunc // Method Not Allowed
}
type namedHandlers struct {
	Created          handlerFunc `sc:"201"`
	NotFound         handlerFunc `sc:"404"`
	MethodNotAllowed handlerFunc `sc:"405"`
}

func (h *handlers) Set(sc int, handler handlerFunc) {
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
	t := reflect.TypeOf(namedHandlers{})
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("sc") == strconv.Itoa(sc) {
			s.FieldByName(t.Field(i).Name).Set(reflect.ValueOf(handler))
			break
		}
	}
}

func initHandlers() *handlers {
	var codes []int
	t := reflect.TypeOf(numHandlers{})
	for i := 0; i < t.NumField(); i++ {
		code, _ := strconv.Atoi(strings.TrimLeft(t.Field(i).Name, "SC"))
		codes = append(codes, code)
	}
	num := numHandlers{}
	nam := namedHandlers{}
	num.init()
	nam.init()
	return &handlers{codes, num, nam}
}
func (h *numHandlers) init() {
	t := reflect.TypeOf(*h)
	s := reflect.ValueOf(h).Elem()
	for i := 0; i < s.NumField(); i++ {
		s.Field(i).Set(reflect.ValueOf(func(SC int, e error) handlerFunc {
			return func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(SC) }
		}(strconv.Atoi(strings.TrimLeft(t.Field(i).Name, "SC")))))
	}
}
func (h *namedHandlers) init() {
	t := reflect.TypeOf(*h)
	s := reflect.ValueOf(h).Elem()
	for i := 0; i < s.NumField(); i++ {
		s.Field(i).Set(reflect.ValueOf(func(SC int, e error) handlerFunc {
			return func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(SC) }
		}(strconv.Atoi(t.Field(i).Tag.Get("sc")))))
	}
}
