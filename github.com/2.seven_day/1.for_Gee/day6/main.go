package main

import (
	"fmt"
	"go_learn/github.com/2.seven_day/1.for_Gee/day6/gee"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlobal("/Users/whp/go/src/go_learn/github.com/2.seven_day/1.for_Gee/day6/templates/*")
	r.Static("/assets", "/Users/whp/go/src/go_learn/github.com/2.seven_day/1.for_Gee/day6/static")

	stu1 := &student{"geektutu", 20}
	stu2 := &student{"jack", 22}

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": []*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2022, 12, 27, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
