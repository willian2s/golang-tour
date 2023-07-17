package more_types

import "fmt"

var m map[string]Vertex

func Map() map[string]Vertex {
	m = make(map[string]Vertex)

	m["apple"] = Vertex{1, 2}

	fmt.Println(m["apple"])

	return m
}

var vertex = map[string]Vertex{
	"apple":  {1, 2},
	"google": {1, 2},
}

func MapLiterals() map[string]Vertex {
	return vertex
}

func MapMutation() {
	m := make(map[string]int)

	m["apple"] = 42
	fmt.Println("The value: ", m["apple"])

	m["apple"] = 48
	fmt.Println("The value: ", m["apple"])

	delete(m, "apple")
	fmt.Println("The value: ", m["apple"])

	v, ok := m["apple"]
	fmt.Println("The value: ", v, "Present: ", ok)
}
