package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 模拟一个读锁重入的情况

var l sync.RWMutex

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan int)

	go func() {
		l.RLock()
		defer l.RUnlock()
		fmt.Println(1)
		c <- 1
		fmt.Println(2)
		runtime.Gosched()
		fmt.Println(3)
		b()
		fmt.Println(4)
		wg.Done()
	}()

	go func() {
		fmt.Println(5)
		<-c
		fmt.Println(6)
		l.Lock()
		fmt.Println(7)
		fmt.Println(8)
		defer l.Unlock()
		fmt.Println(9)
		wg.Done()
	}()

	go func() {
		i := 1
		for {
			i++
		}
	}()
	wg.Wait()
}

func b(){
	fmt.Println(10)
	l.RLock()
	fmt.Println(11)
	defer l.RUnlock()
	fmt.Println(12)
}

/*
加锁的经验：
	1.运行时离开当前逻辑就释放锁
	2.锁的粒度越小越好，加锁后尽快释放锁
	3.尽量不用 defer 释放锁
	4.读锁不要嵌套
*/
