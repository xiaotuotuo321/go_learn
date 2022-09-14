package main

import "fmt"

func main() {
	arrayA := [2]int{100, 200}
	testArrayPoint(&arrayA)

	arrayB := arrayA[:]
	testArrayPoint2(arrayB)
	fmt.Println(arrayA)
}

func testArrayPoint(x *[2]int) {
	fmt.Printf("func Array: %p, %v\n", x, *x)
	(*x)[1] += 100
}

func testArrayPoint2(x []int) {
	fmt.Printf("func Array: %p, %v", &x, x)
	x[1] += 100
}
