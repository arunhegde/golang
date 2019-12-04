package main

import (
	"fmt"
)

type day string

func (d day) updateDayFails() {
	d = "Monday"
}

func (d day) printDay() {
	fmt.Println(d)
}

func main() {
	var d day = "Sunday"
	d.updateDayFails()
	d.printDay()
}