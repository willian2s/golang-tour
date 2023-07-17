package more_types

import "fmt"

type Vertex struct {
	X int
	Y int
}

// Struct returns a Vertex with the values (1, 2).
//
// This function does not take any parameters.
// It returns a Vertex.
func Struct(x, y int) Vertex {
	return Vertex{x, y}
}

func StructField(x, y int) Vertex {
	v := Vertex{1, 2}
	v.X = 3

	return v
}

func StructPointer(x, y int) Vertex {
	v := Vertex{1, 2}
	p := &v

	p.X = 1e9

	return v
}

var (
	v1 = Vertex{1, 2}       // has type Vertex
	v2 = Vertex{X: 1, Y: 0} // Y:0 is implicit
	v3 = Vertex{}           // X:0 and Y:0
	p  = &Vertex{1, 2}      // has type *Vertex
)

func StrucLiteral() {
	fmt.Println(v1, v2, v3, p)
}
