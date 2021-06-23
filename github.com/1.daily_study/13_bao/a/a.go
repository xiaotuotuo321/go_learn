package a

import (
	"fmt"
	"go_learn/github.com/1.daily_study/13_bao/b"
)

func init(){
	fmt.Println("我是a包中的init")
}

func A() {
	b.B()
	fmt.Println("A方法")
}
