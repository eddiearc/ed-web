package edw

import (
	"html/template"
	"net/http"
	"path"
)

type RouterGroup struct {
	prefix        string
	middlewares   []HandlerFunc
	engine        *Engine
	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
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

func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

// Static serve static files
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	group.GET(urlPattern, handler)
}
