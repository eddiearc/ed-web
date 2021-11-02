package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(writer, "Welcome to ed-web.")
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(writer, "Other url request = %q\n", req.URL.Path)
	}
}

func main() {
	engine := &Engine{}
	log.Fatal(http.ListenAndServe(":9999", engine))
}
