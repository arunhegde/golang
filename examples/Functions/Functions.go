package main

import "fmt"

func AddTwoNumbers(a, b int32) int32 {
	return a+b
}

// Functionals are first-class citizens in Golang
// Functions can be assigned to a valiraibles.
func main() {
	//Assign a function to a variable
	var b func(int32, int32) int32
	b= AddTwoNumbers
	fmt.Println("Sum: ", b(4,6))

	// An anonymous function is assigned to a variable
	a := func() {
		fmt.Println("hello world first class function")
	}

	a()
	fmt.Printf("%T\n", a)

	//Pass function as arguments
	fmt.Println("Sum of three numbers 5, 6, 9 is: ", AddThreeNumbers(5, 6, 9, AddTwoNumbers))
}

func AddThreeNumbers(a, b, c int32, addTwo func(int32, int32) int32) int32 {
	return a + addTwo(b, c)
}
