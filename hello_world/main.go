package main

import (
	"fmt"
	"time"
)

// main is the entry point of the Go program.
//
// It does not take any parameters.
// It does not return any values.
func main() {
	greet := helloworld("Willian")

	fmt.Println(greet)
	fmt.Println(time.Now())
}

// helloworld returns a formatted string that says hello to the given name.
//
// It takes a parameter:
// - name: a string representing the name of the person to greet.
//
// It returns a string.
func helloworld(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
