package main

import (
	"fmt"
	"runtime"
	"time"

	flowControl "github.com/willian2s/golang-tour/basic/flow_control"
	"github.com/willian2s/golang-tour/basic/packages"
)

func mainPackages() {
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
	packages.BasicTypes()
	packages.Zero()
	packages.TypeConversion()
	packages.TypeInference()
	packages.Constants()
	packages.ConstantsNumeric()
}

func mainFlowControl() {
	forSum := flowControl.For()
	fmt.Println("For: ", forSum)
	forSumContinue := flowControl.For()
	fmt.Println("ForContinue: ", forSumContinue)

	while := flowControl.While()
	fmt.Println("While: ", while)

	sqrtPositive := flowControl.Sqrt(2)
	fmt.Println("If: ", sqrtPositive)
	sqrtNegative := flowControl.Sqrt(-4)
	fmt.Println("If: ", sqrtNegative)

	powIfOne := flowControl.Pow(3, 2, 10)
	fmt.Println("Pow: ", powIfOne)
	powIfTwo := flowControl.Pow(3, 3, 20)
	fmt.Println("Pow: ", powIfTwo)

	powIfElseOne := flowControl.Pow(3, 2, 10)
	fmt.Println("Pow: ", powIfElseOne)
	powIfElseTwo := flowControl.Pow(3, 3, 20)
	fmt.Println("Pow: ", powIfElseTwo)

	exerciseLoops := flowControl.SqrtExercise(2)
	fmt.Println("ExerciseLoops: ", exerciseLoops)

	switchOs := flowControl.Switch(runtime.GOOS)
	fmt.Println("SwitchOs: ", switchOs)

	isSaturday := flowControl.SwitchEvaluationOrder(time.Now().Weekday())
	fmt.Println("SwitchEvaluationOrder: ", isSaturday)

	time := flowControl.SwitchWithNoCondition()
	fmt.Println("SwitchWithNoCondition: ", time)

	flowControl.Defer()
	flowControl.DeferMulti() // LIFO = Last In First Out
}

func main() {
	mainPackages()

	mainFlowControl()
}
