package main

import (
	"fmt"
	"math"

	"github.com/willian2s/golang-tour/methods_interfaces/packages"
)

func main() {
	// a := packages.Method{X: 3, Y: 4}
	// fmt.Println(a.Abs())

	// b := packages.MyFloat(3.4)
	// fmt.Println(b.Abs())

	// c := packages.Method{X: 3, Y: 4}
	// c.Scale(10)
	// fmt.Println(c.Abs())

	var abser packages.Abser
	float := packages.MyFloat(-math.Sqrt2)
	vertex := packages.Method{X: 3, Y: 4}

	abser = float
	abser = &vertex

	abser = vertex

	fmt.Println(abser.Abs())

	var inter packages.I

	var null *packages.T
	inter = null
	packages.Describe(inter)
	inter.M()

	inter = &packages.T{S: "Hello"}
	packages.Describe(inter)
	inter.M()
}
