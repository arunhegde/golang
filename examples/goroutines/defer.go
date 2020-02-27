package main

import (
	"fmt"
)

func main() {
	aValue := new(int)
	*aValue = 1

	//Defer the following code until all the functions (in this case main routine) is completed
	defer fmt.Println("aValue = ", *aValue)

	for i :=0; i<100; i++ {
		*aValue++
	}

}
