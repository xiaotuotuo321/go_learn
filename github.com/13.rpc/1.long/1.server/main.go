package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

type Rect struct {
}

// Area RPC方法 求矩形的面积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Height * p.Width
	return nil
}

func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	return nil
}

func main() {
	rect := new(Rect)
	rpc.Register(rect)

	rpc.HandleHTTP()
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Panicln(err)
	}
}
