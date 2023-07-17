package flow_control

import "fmt"

func Defer() {
	defer fmt.Println("Goodbye")

	fmt.Print("Hello\n")
}
