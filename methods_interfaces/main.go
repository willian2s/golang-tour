package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"github.com/willian2s/golang-tour/methods_interfaces/packages"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
)

func main() {
	a := packages.Method{X: 3, Y: 4}
	fmt.Println(a.Abs())

	b := packages.MyFloat(3.4)
	fmt.Println(b.Abs())

	c := packages.Method{X: 3, Y: 4}
	c.Scale(10)
	fmt.Println(c.Abs())

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

	packages.TypeSwitch(21)
	packages.TypeSwitch("hello")
	packages.TypeSwitch(true)

	personOne := packages.Person{
		Name: "Willian",
		Age:  26,
	}
	personTwo := packages.Person{
		Name: "Santos",
		Age:  27,
	}
	fmt.Println(personOne, personTwo)

	// Exercise Stringer
	hosts := map[string]packages.IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	// End of Exercise

	if err := packages.Run(); err != nil {
		fmt.Println(err)
	}

	// Exercise Error
	packages.RunCode()
	// End of Exercise

	packages.Reader("inconstitucionalissimamente")

	// Exercise Reader
	reader.Validate(packages.AReader{})
	// End of Exercise

	// Exercise ROT13 Reader
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := packages.Rot13Reader{R: s}
	io.Copy(os.Stdout, &r)
	// End of Exercise

	// Exercise Image
	m := packages.Image{Width: 256, Height: 256}
	pic.ShowImage(m)
	// End of Exercise
}
