package b

import "fmt"

func init() {
	fmt.Println("我是b包中的init")
}

func B() {
	fmt.Println("B方法")
}