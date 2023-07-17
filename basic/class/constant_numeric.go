package class

import "fmt"

const (
	Big = 1 << 100

	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 { return x * 0.1 }

func ConstantsNumeric() {
	fmt.Printf("Need Int: %d\n", needInt(Small))
	fmt.Printf("Need Float: %f\n", needFloat(Small))
	fmt.Printf("Need Int: %f\n", needFloat(Big))
}
