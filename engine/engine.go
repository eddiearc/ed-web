package engine

import (
	"fmt"
	"net/http"
)

type (
	HandlerFunc func(writer http.ResponseWriter, req *http.Request)
	Method      string
	Pattern     string
)

type Engine struct {
	router map[Method]map[Pattern]HandlerFunc
}

func (engine *Engine) addRouter(method Method, pattern string, handlerFunc HandlerFunc) {
	if engine.router[method] == nil {
		engine.router[method] = make(map[Pattern]HandlerFunc)
	}
	engine.router[method][Pattern(pattern)] = handlerFunc
}

func (engine *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	engine.addRouter("GET", pattern, handlerFunc)
}

func (engine *Engine) POST(pattern string, handlerFunc HandlerFunc) {
	engine.addRouter("POST", pattern, handlerFunc)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func New() *Engine {
	return &Engine{
		make(map[Method]map[Pattern]HandlerFunc),
	}
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	if handlerFunc, found := engine.router[Method(req.Method)][Pattern(req.URL.Path)]; found {
		handlerFunc(writer, req)
	} else {
		_, _ = fmt.Fprintf(writer, "404 NOT FOUND.")
	}
}
