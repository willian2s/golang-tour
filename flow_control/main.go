package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/willian2s/golang-tour/flow_control/class"
)

func main() {
	forSum := class.For()
	fmt.Println("For: ", forSum)
	forSumContinue := class.For()
	fmt.Println("ForContinue: ", forSumContinue)

	while := class.While()
	fmt.Println("While: ", while)

	sqrtPositive := class.Sqrt(2)
	fmt.Println("If: ", sqrtPositive)
	sqrtNegative := class.Sqrt(-4)
	fmt.Println("If: ", sqrtNegative)

	powIfOne := class.Pow(3, 2, 10)
	fmt.Println("Pow: ", powIfOne)
	powIfTwo := class.Pow(3, 3, 20)
	fmt.Println("Pow: ", powIfTwo)

	powIfElseOne := class.Pow(3, 2, 10)
	fmt.Println("Pow: ", powIfElseOne)
	powIfElseTwo := class.Pow(3, 3, 20)
	fmt.Println("Pow: ", powIfElseTwo)

	exerciseLoops := class.SqrtExercise(2)
	fmt.Println("ExerciseLoops: ", exerciseLoops)

	switchOs := class.Switch(runtime.GOOS)
	fmt.Println("SwitchOs: ", switchOs)

	isSaturday := class.SwitchEvaluationOrder(time.Now().Weekday())
	fmt.Println("SwitchEvaluationOrder: ", isSaturday)

	time := class.SwitchWithNoCondition()
	fmt.Println("SwitchWithNoCondition: ", time)

	class.Defer()
	class.DeferMulti() // LIFO = Last In First Out
}
