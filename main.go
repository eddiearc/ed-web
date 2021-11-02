package main

import (
	"ed-web/engine"
	"fmt"
	"net/http"
)

func main() {
	e := engine.New()
	e.GET("/", func(writer http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(writer, "Welcome to ed-web.")
	})
	_ = e.Run(":9999")
}
