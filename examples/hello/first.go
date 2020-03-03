package main

import (
	"fmt"
	"github.com/arunhegde/golang/go-modules-test/mypackage"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args[1:])
	fmt.Println( add(os.Args[1], os.Args[2]))
	mypackage.HelloFunction()
}

func add(a, b string) int {
	arg1, _ :=  strconv.Atoi(a)
	arg2, _ := strconv.Atoi(b)

	return arg1 + arg2
}