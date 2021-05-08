package gee

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc){
	e.router.addRoute(method, pattern, handler)
}

func (e *Engine) Get(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) Post(pattern string, handler HandlerFunc){
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) ServeHTTP (w http.ResponseWriter, req *http.Request){
	c := NewContext(w, req)
	e.router.handle(c)
}

func (e *Engine) Run(addr string) (err error){
	return http.ListenAndServe(addr, e)
}