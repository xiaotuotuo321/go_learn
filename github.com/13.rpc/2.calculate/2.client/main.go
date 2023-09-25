package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	A int
	B int
}

func main() {
	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	ret := 0
	err2 := conn.Call("Rect.Multiplier", Params{10, 5}, &ret)
	if err2 != nil {
		log.Panicln(err2)
	}
	fmt.Println("乘法: ", ret)

	err3 := conn.Call("Rect.Division", Params{10, 5}, &ret)
	if err3 != nil {
		log.Panicln(err3)
	}
	fmt.Println("除法：", ret)
}
