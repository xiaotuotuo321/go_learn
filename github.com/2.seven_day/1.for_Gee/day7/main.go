package main

import (
	"go_learn/github.com/2.seven_day/1.for_Gee/day7/gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "hello hello\n")
	})

	r.GET("/panic", func(c *gee.Context) {
		names := []string{"xiaoming"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
