package main

import (
	"fmt"

	"github.com/willian2s/golang-tour/methods_interfaces/packages"
)

func main() {
	a := packages.Method{X: 3, Y: 4}
	fmt.Println(a.MAbs())

	b := packages.MyFloat(3.4)
	fmt.Println(b.FAbs())

	c := packages.Method{X: 3, Y: 4}
	c.Scale(10)
	fmt.Println(c.MAbs())
}
