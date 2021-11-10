package edw

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	engine      *Engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	e := group.engine
	newGroup := &RouterGroup{
		prefix:      group.prefix + prefix,
		middlewares: []HandlerFunc{},
		engine:      e,
	}
	e.groups = append(e.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) groupAddRouter(method Method, pattern string, handler HandlerFunc) {
	urlPath := group.prefix + pattern
	group.engine.addRouter(method, urlPath, handler)
}

func (group *RouterGroup) REQUEST(pattern string, handler HandlerFunc) {
	group.groupAddRouter(REQUEST, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.groupAddRouter(GET, pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.groupAddRouter(POST, pattern, handler)
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}
