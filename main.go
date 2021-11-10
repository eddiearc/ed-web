package main

import (
	"ed-web/edw"
	"net/http"
)

func main() {
	e := edw.New(func(c *edw.Context) {
		c.String(http.StatusOK, "Welcome to ed-web.")
	})

	e.Use(edw.Logger())

	_ = e.Run(":9999")
}
