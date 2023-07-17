package packages

import (
	"fmt"
	"math/rand"
)

// FunctionContinue generates and returns a string representing a random number.
//
// This function takes no parameters.
// It returns a string.
func FunctionContinue() string {
	return fmt.Sprintf("My favorite number is: %d", rand.Intn(1000))
}
