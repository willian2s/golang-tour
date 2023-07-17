package more_types

import "fmt"

// Pointer is a function that demonstrates the use of pointers in Go.
//
// It takes two integer parameters, i and j.
// It returns two integers.
func Pointer(i, j int) (int, int) {

	p := &i // pointer to i
	fmt.Println(*p)
	*p = 21 // set i to 21
	fmt.Println(i)

	p = &j // pointer to j
	fmt.Println(*p)
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j)

	return i, j
}
