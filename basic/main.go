package main

import (
	"fmt"

	"github.com/willian2s/golang-tour/basic/packages"
)

func main() {
	pi := packages.ImportNamed()
	fmt.Println(pi)

	sum := packages.Function(1, 2)
	fmt.Println(sum)

	myFavNumber := packages.FunctionContinue()
	fmt.Println(myFavNumber)

	swapOne, swapTwo := packages.MultipleResults("Willian", "Santos")
	fmt.Println(swapOne, swapTwo)

	x, y := packages.NamedResults(17)
	fmt.Println(x, y)

	packages.Variables()
	packages.VariablesWIthInitializers()
	packages.SortVariableDeclaration()
}
