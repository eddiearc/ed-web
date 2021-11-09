package engine

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	engine      *Engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	return &RouterGroup{
		prefix:      group.prefix + prefix,
		middlewares: make([]HandlerFunc, 0),
		engine:      group.engine,
	}
}

func (group *RouterGroup) addRouter(method Method, pattern string, handler HandlerFunc) {
	urlPath := group.prefix + pattern
	group.engine.addRouter(method, urlPath, handler)
}

func (group *RouterGroup) REQUEST(pattern string, handler HandlerFunc) {
	group.addRouter(REQUEST, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRouter(GET, pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRouter(POST, pattern, handler)
}
