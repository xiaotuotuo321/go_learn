package main

import (
	"fmt"
	"go_learn/github.com/seven_day/for_Gee/day1/3.base_3/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Get("/", func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintf(w, "URL.Path = %q \n", req.URL.Path)
	})

	r.Get("/hello", func(w http.ResponseWriter, req *http.Request){
		for k, v := range req.Header{
			fmt.Fprintf(w, "Header[%q] = %q \n", k, v)
		}
	})

	r.Run(":9999")
}


