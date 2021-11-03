package engine

import (
	"net/http"
)

type router struct {
	// handlers handle http request.
	handlers map[string]map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]map[string]HandlerFunc),
	}
}

func (r *router) addRouter(method string, pattern string, handlerFunc HandlerFunc) {
	if r.handlers[method] == nil {
		r.handlers[method] = make(map[string]HandlerFunc)
	}
	r.handlers[method][pattern] = handlerFunc
}

func (r *router) handle(c *Context) {
	if handlerFunc, found := r.handlers[c.Method][c.Path]; found {
		handlerFunc(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND. %s\n", c.Path)
	}
}
