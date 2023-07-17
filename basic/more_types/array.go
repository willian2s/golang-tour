package more_types

import (
	"errors"
	"fmt"
	"strings"
)

func Array() ([2]string, [6]int) {
	var array [2]string

	array[0] = "Hello"
	array[1] = "World"

	primes := [6]int{2, 3, 5, 7, 11, 13}

	return array, primes
}

func Slice() ([6]int, []int) {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var slice []int = primes[1:4]

	return primes, slice
}

func SlicePointers() [4]string {
	name := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(name)

	a := name[0:2]
	b := name[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(name)

	return name
}

type sliceLiteral struct {
	i int
	b bool
}

func SliceLiterals() ([]int, []bool, []sliceLiteral) {
	q := []int{1, 2, 3, 4, 5}
	fmt.Println(q)

	r := []bool{true, false, true, true, false}
	fmt.Println(r)

	s := []sliceLiteral{
		{2, true},
		{3, false},
		{4, true},
		{5, false},
	}

	return q, r, s
}

func SliceBounds() []int {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	return s
}

func SliceLenCap() string {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	return fmt.Sprintf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func SliceNil(slice []int) ([]int, error) {
	s := slice
	if s == nil {
		return nil, errors.New("slice is nil")
	}

	return s, nil
}

func MakingSlice() {
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)
}

func SliceOfSlice() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	board[0][1] = "O"
	board[2][0] = "X"
	board[2][1] = "O"
	board[1][1] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func Append() string {
	var s []int

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

	return fmt.Sprintf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
