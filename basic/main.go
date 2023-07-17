package main

import (
	"fmt"
	"runtime"
	"time"

	flowControl "github.com/willian2s/golang-tour/basic/flow_control"
	moreTypes "github.com/willian2s/golang-tour/basic/more_types"
	"github.com/willian2s/golang-tour/basic/packages"
	"golang.org/x/tour/pic"
)

func MainPackages() {
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

func MainFlowControl() {
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

func MainMoreTypes() {
	pointerOne, pointerTwo := moreTypes.Pointer(1, 2)
	fmt.Println("Pointer", pointerOne, pointerTwo)

	vertex := moreTypes.Struct(1, 2)
	fmt.Println("Struct", vertex)

	vertex = moreTypes.StructField(1, 2)
	fmt.Println("StructField", vertex)

	vertex = moreTypes.StructPointer(1, 2)
	fmt.Println("StructPointer", vertex)

	moreTypes.StrucLiteral()

	arrayOne, arrayTwo := moreTypes.Array()
	fmt.Println("Array", arrayOne, arrayTwo)

	array, slice := moreTypes.Slice()
	fmt.Println("Array", array, "Slice", slice)

	name := moreTypes.SlicePointers()
	fmt.Println("Name", name)

	sliceInteger, sliceBoolean, sliceStruc := moreTypes.SliceLiterals()
	fmt.Println("SliceInteger", sliceInteger, "SliceBoolean", sliceBoolean, "SliceStruc", sliceStruc)

	sliceBounds := moreTypes.SliceBounds()
	fmt.Println("SliceBounds", sliceBounds)

	sliceLenCaps := moreTypes.SliceLenCap()
	fmt.Println("SliceLenCap", sliceLenCaps)

	var mySlice []int
	sliceNil, err := moreTypes.SliceNil(mySlice)
	fmt.Println("SliceNil", sliceNil, "Error", err)

	mySlice = []int{1, 2, 3, 4, 5}
	sliceNil, err = moreTypes.SliceNil(mySlice)
	fmt.Println("SliceNil", sliceNil, "Error", err)

	moreTypes.MakingSlice()
	moreTypes.SliceOfSlice()

	sliceAppend := moreTypes.Append()
	fmt.Println("SliceAppend", sliceAppend)

	moreTypes.Range()
	moreTypes.MakeRange()

	// Slice Exercise
	pic.Show(moreTypes.Pic)
	// Slice Exercise
}

func main() {
	MainMoreTypes()
}
