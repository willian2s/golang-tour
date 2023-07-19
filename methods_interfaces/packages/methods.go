package packages

import "math"

type Method struct {
	X, Y float64
}

type MyFloat float64

type Abser interface {
	Abs() float64
}

func (m Method) Abs() float64 {
	return math.Sqrt(m.X*m.X + m.Y*m.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (m *Method) Scale(f float64) {
	m.X = m.X * f
	m.Y = m.Y * f
}
