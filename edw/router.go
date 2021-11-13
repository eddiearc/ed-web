package edw

import (
	"net/http"
	"strings"
)

type Router struct {
	// root route and match request that execute Handler.
	root *Node
}

type Method int8

const (
	REQUEST Method = 1 + iota
	GET
	POST
)

func (r *Router) addRouter(method Method, pattern string, handler HandlerFunc) {
	urlPath, err := formatUrlPath(pattern)
	if err != nil {
		panic(err)
	}

	parts := getParts(urlPath)

	r.root.insert(method, urlPath, parts, handler, 0)
	if err != nil {
		panic(err)
	}
}

func (r *Router) handle(c *Context) {
	var method Method
	switch c.Method {
	case "GET":
		method = GET
	case "POST":
		method = POST
	}

	urlPath, _ := formatUrlPath(c.Path)
	parts := getParts(urlPath)

	node := r.root.search(method, urlPath, parts, 0)
	if node != nil {
		if node.Wildcard {
			r.wildcardParameter(node.Pattern, parts, c)
		}
		c.handlers = append(c.handlers, node.Handler)
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND.\n PATH: %s\n", c.Path)
		})
	}

	c.Next()
}

func (r *Router) wildcardParameter(pattern string, parts []string, c *Context) {
	patternParts := getParts(pattern)

	for i, p := range patternParts {
		if i >= len(parts) {
			return
		}
		// restful match parameter.
		if p[0] == ':' && len(p) > 1 {
			c.Params[p[1:]] = parts[i]
		}
		// fuzzy match parameter.
		if p[0] == '*' && len(p) > 1 {
			c.Params[p[1:]] = strings.Join(parts[i:], "/")
			break
		}
	}
}
