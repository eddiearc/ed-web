package main

import (
	"ed-web/edw"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	e := edw.Default()

	e.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})

	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./static")

	member1 := &struct {
		Name string
		Age  int
	}{Name: "eddie", Age: 20}

	member2 := &struct {
		Name string
		Age  int
	}{Name: "jack", Age: 22}

	e.GET("/", func(c *edw.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	e.GET("/members", func(c *edw.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", edw.JSON{
			"title": "ed-web",
			"stuArr": [2]*struct {
				Name string
				Age  int
			}{member1, member2},
		})
	})

	e.GET("/panic", func(c *edw.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	e.GET("/date", func(c *edw.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", edw.JSON{
			"title": "ed-web",
			"now":   time.Date(2021, 11, 10, 20, 0, 0, 0, time.UTC),
		})
	})

	_ = e.Run(":9999")
}
