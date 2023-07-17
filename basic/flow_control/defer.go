package class

import "fmt"

func Defer() {
	defer fmt.Println("Goodbye")

	fmt.Print("Hello\n")
}
