package packages

// NamedResults splits an integer into two parts.
//
// It takes an integer as input and returns two integers as output.
func NamedResults(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
