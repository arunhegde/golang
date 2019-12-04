package main

import (
	"fmt"
)

func updateSlice(d []string) {
	d[0] = "Funday"
}

type myDay string

func updateDay(d *myDay) {
	*d = "XYZDay"
}

func main() {
	days := []string{"Sunday", "Monday", "Tuesday"}
	updateSlice(days)
	fmt.Println(days)

	d := myDay("Wednesday")
	updateDay(&d)
	fmt.Println(d)
}