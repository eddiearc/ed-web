package main

import (
	"ed-web/edw"
	"log"
	"net/http"
	"time"
)

func onlyForV2() edw.HandlerFunc {
	return func(c *edw.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.String(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	e := edw.New(func(c *edw.Context) {
		c.String(http.StatusOK, "Welcome to ed-web.")
	})

	e.Use(edw.Logger())
	v2 := e.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *edw.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	_ = e.Run(":9999")
}
