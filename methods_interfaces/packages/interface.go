package packages

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}

	fmt.Println(t.S)
}

func (f F) M() {
	fmt.Println(f)
}

func Describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
