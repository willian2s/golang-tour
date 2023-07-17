package class

// For calculates the sum of numbers from 0 to 9 using a for loop.
//
// No parameters.
// Returns the sum of the numbers.
func For() int {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	return sum
}

func ForContinue() int {
	sum := 1

	for sum < 10 {
		sum += sum
	}

	return sum
}
