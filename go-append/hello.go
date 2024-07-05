package main

import "fmt"

func main() {
	fmt.Printf("Test 1:\n")
	test1()
	fmt.Printf("\nTest 2:\n")
	test2()
}


func test1() {
	// create an array
	a := []int{}
	a = append(a, 1, 2, 3, 4, 5, 6, 7)

	// append to it
	b := append(a, 10) // will this copy the array or not ?

	// modify an already existing element
	a[0] = 100

	// ahah funny behaviour
	fmt.Printf("%v\n", a) // [100 2 3 4 5 6 7]
	fmt.Printf("%v\n", b) // [100 2 3 4 5 6 7 10]
}

// let's do another test shall we ?
func test2() {

	// let's create the same array, *but in a different way*
	a := []int{1, 2, 3, 4, 5, 6, 7}

	// append to it
	b := append(a, 10) // will this copy the array or not ?

	// modify an already existing element
	a[0] = 100

	// ahah different behaviour (standard copy behaviour)
	fmt.Printf("%v\n", a) // [100 2 3 4 5 6 7]
	fmt.Printf("%v\n", b) // [1 2 3 4 5 6 7 10]
}

// TLDR; append() sometimes copies the array, and sometimes doesn't
// the Go notion of array and slice is important here, look it up