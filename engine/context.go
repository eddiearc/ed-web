package engine

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Context Http request context.
type Context struct {
	// about http connect.
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path       string
	Method     string
	StatusCode int
	Params     map[string]string
	handlers   []HandlerFunc
	index      int
}

func newContext(writer http.ResponseWriter, req *http.Request, middlewares []HandlerFunc) *Context {
	return &Context{
		Writer:   writer,
		Req:      req,
		Path:     req.URL.Path,
		Method:   req.Method,
		Params:   map[string]string{},
		handlers: middlewares,
		index:    -1,
	}
}

// Param restful param.
func (c *Context) Param(key string) string {
	return c.Params[key]
}

// Query get method value in url.
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// PostForm post method value in post form.
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Status set status.
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader set k-v pair in http header.
func (c *Context) SetHeader(k string, v string) {
	c.Writer.Header().Set(k, v)
}

// JSON http response json format data.
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// HTML http response html format data.
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if _, err := c.Writer.Write([]byte(html)); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data http response data.
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	if _, err := c.Writer.Write(data); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// String format string message.
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	_, _ = c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// Next execute middleware.
func (c *Context) Next() {
	c.index++
	if c.index < len(c.handlers) {
		c.handlers[c.index](c)
	} else {
		return
	}
}
