package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type regexResolver struct {
	handlers map[string]http.HandlerFunc
	cache    map[string]*regexp.Regexp
}

func newPathResolver() *regexResolver {
	return &regexResolver{
		handlers: make(map[string]http.HandlerFunc),
		cache:    make(map[string]*regexp.Regexp),
	}
}

func (r *regexResolver) Add(regex string, handler http.HandlerFunc) {
	r.handlers[regex] = handler
	cache, _ := regexp.Compile(regex)
	r.cache[regex] = cache
}

func (r *regexResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	for pattern, handlerFunc := range r.handlers {
		if r.cache[pattern].MatchString(check) {
			handlerFunc(res, req)
			return
		}
	}
	http.NotFound(res, req)
}

func main() {
	rr := newPathResolver()
	rr.Add("GET /hello", hello)
	rr.Add("(GET|HEAD) /goodbye(/?[A-Za-z0-9]*)?", goobye)
	rr.Add("GET /", homePage)
	http.ListenAndServe(":8080", rr)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "default name"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func goobye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := ""
	if len(parts) > 2 {
		name = parts[2]
	}
	if name == "" {
		name = "default name"
	}
	fmt.Fprint(res, "Goodbye ", name)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "My homepage")
}
