package exinit

import (
	"fmt"
)

var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func init() {
	d++
	fmt.Println("init", d)
}

func f() int {
	d++
	fmt.Println("f()", d)
	return d
}

func PrintD() {
	fmt.Println("Print", d)
}
