package main

import (
	"ed-web/engine"
	"net/http"
)

func main() {
	e := engine.New(func(c *engine.Context) {
		c.String(http.StatusOK, "Welcome to ed-web.")
	})

	_ = e.Run(":9999")
}
