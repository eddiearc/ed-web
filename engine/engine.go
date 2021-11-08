package engine

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router *Router
}

func (engine *Engine) addRouter(method Method, pattern string, handlerFunc HandlerFunc) {
	engine.router.addRouter(method, pattern, handlerFunc)
}

func (engine *Engine) REQUEST(pattern string, handler HandlerFunc) {
	engine.addRouter(REQUEST, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRouter(GET, pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRouter(POST, pattern, handler)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func New(rootHandler HandlerFunc) *Engine {
	return &Engine{
		newRouter(rootHandler),
	}
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	engine.router.handle(newContext(writer, req))
}
