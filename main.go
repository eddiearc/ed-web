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
	e := edw.New(func(c *edw.Context) {
		c.String(http.StatusOK, "Welcome to ed-web.")
	})
	e.Use(edw.Logger())
	e.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})

	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./static")

	stu1 := &struct {
		Name string
		Age  int
	}{Name: "eddie", Age: 20}

	stu2 := &struct {
		Name string
		Age  int
	}{Name: "jack", Age: 22}

	e.GET("/", func(c *edw.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	e.GET("/students", func(c *edw.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", edw.JSON{
			"title": "gee",
			"stuArr": [2]*struct {
				Name string
				Age  int
			}{stu1, stu2},
		})
	})

	e.GET("/date", func(c *edw.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", edw.JSON{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	_ = e.Run(":9999")
}
