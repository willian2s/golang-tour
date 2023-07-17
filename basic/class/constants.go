package class

import (
	"fmt"
	"math"
)

const Pi = math.Pi

func Constants() {
	const Word = "Willian"
	fmt.Printf("Hello, %s!\n", Word)
	fmt.Printf("Happy %f day!\n", Pi)

	const Truth = true
	fmt.Printf("Go Rules? %v\n", Truth)
}
