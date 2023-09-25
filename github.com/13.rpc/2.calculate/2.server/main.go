package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Params struct {
	A, B int
}

type Rect struct {
}

func (r *Rect) Multiplier(p *Params, ret *int) error {
	*ret = p.B * p.A
	return nil
}

func (r *Rect) Division(p *Params, ret *int) error {
	if p.B == 0 {
		*ret = 0
		return nil
	}
	*ret = p.A / p.B

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
