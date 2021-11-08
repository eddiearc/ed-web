package engine

import (
	"net/http"
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

func newRouter(rootHandler HandlerFunc) *Router {
	return &Router{
		root: newTrieRoot(rootHandler),
	}
}

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
			// TODO parse Wildcard k-v, store in context.
		}
		node.Handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND.\n PATH: %s\n", c.Path)
	}
}
