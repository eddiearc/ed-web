package engine

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	*RouterGroup
	router *Router
	groups []*RouterGroup // store all groups
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

func New(rootHandler HandlerFunc) (e *Engine) {
	e = &Engine{
		router: newRouter(rootHandler),
	}
	e.groups = []*RouterGroup{e.RouterGroup}
	e.RouterGroup = &RouterGroup{engine: e}
	return e
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	engine.router.handle(newContext(writer, req))
}
