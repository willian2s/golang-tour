package more_types

func Compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
