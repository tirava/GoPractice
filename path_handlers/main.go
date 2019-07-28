package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	//fmt.Fprint(res, "qqq")
	for pattern, handlerFunc := range p.handlers {
		if ok, err := path.Match(pattern, check); ok && err == nil {
			handlerFunc(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, req)
		}
	}
	//fmt.Fprint(res, "qqq")
	http.NotFound(res, req)
}

func main() {
	pr := newPathResolver()
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye/*", goobye)
	pr.Add("GET /", homePage)
	http.ListenAndServe(":8080", pr)
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
	name := parts[2]
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
