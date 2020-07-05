package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func new_router() *Router {
	return &Router{
		rules : make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) findHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExists := r.rules[path][method]
	return handler, methodExists, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExists, exist := r.findHandler(request.URL.Path, request.Method)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return 
	}

	if !methodExists {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)

}