package snow

import (
	"fmt"
	"go_learn/github.com/daily_study/13_bao/calc"
)

var (
	a, b, c, d int
)

const (
	x = 10
	y = 2
)

func Snow() {
	a = calc.Add(x, y)
	b = calc.Sub(x, y)
	c = calc.Mult(x, y)
	d = calc.Except(x, y)
	fmt.Printf("%d,%d的和为：%d，差为：%d，乘积：%d，商为：%d \n", x, y, a, b, c, d)
}
