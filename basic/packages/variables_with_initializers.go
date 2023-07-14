package packages

import "fmt"

var i, j int = 1, 2

func VariablesWIthInitializers() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
