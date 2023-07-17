package class

import "fmt"

func TypeInference() {
	a := 42
	fmt.Printf("Type: %T, value: %v\n", a, a)
	b := 42.0
	fmt.Printf("Type: %T, value: %v\n", b, b)
	c := 0.867 + 0.5i
	fmt.Printf("Type: %T, value: %v\n", c, c)
}
