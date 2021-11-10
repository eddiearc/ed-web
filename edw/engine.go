package edw

import (
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

type Engine struct {
	*RouterGroup
	*Router
	groups []*RouterGroup // store all groups
}

func (engine *Engine) REQUEST(pattern string, handler HandlerFunc) {
	engine.addRouter(REQUEST, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.Router.addRouter(GET, pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.Router.addRouter(POST, pattern, handler)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func New(rootHandler HandlerFunc) (e *Engine) {
	e = &Engine{
		// router main page.
		Router: &Router{
			root: newTrieRoot(rootHandler),
		},
	}
	e.RouterGroup = &RouterGroup{engine: e}
	e.groups = []*RouterGroup{e.RouterGroup}
	return e
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc

	// check current url match those group prefix and apply (middlewares).
	for _, g := range engine.groups {
		if strings.HasPrefix(req.URL.Path, g.prefix) {
			middlewares = append(middlewares, g.middlewares...)
		}
	}

	engine.handle(newContext(writer, req, middlewares))
}
