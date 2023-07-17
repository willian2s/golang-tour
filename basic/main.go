package main

import (
	"fmt"

	"github.com/willian2s/golang-tour/basic/class"
)

func main() {
	pi := class.ImportNamed()
	fmt.Println(pi)

	sum := class.Function(1, 2)
	fmt.Println(sum)

	myFavNumber := class.FunctionContinue()
	fmt.Println(myFavNumber)

	swapOne, swapTwo := class.MultipleResults("Willian", "Santos")
	fmt.Println(swapOne, swapTwo)

	x, y := class.NamedResults(17)
	fmt.Println(x, y)

	class.Variables()
	class.VariablesWIthInitializers()
	class.SortVariableDeclaration()
	class.BasicTypes()
	class.Zero()
	class.TypeConversion()
	class.TypeInference()
	class.Constants()
	class.ConstantsNumeric()
}
