package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
	threadName string
}

func outputText(j *Job) {
	fileName := j.text + ".txt"
	fileContents := ""
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fileContents += j.text
		fmt.Println(j.text + " : " + j.threadName)
		j.i++
	}

	err := ioutil.WriteFile(fileName, []byte(fileContents), 0644)
	if (err != nil) {
		panic("Something went awry")
	}
}
func main() {
	hello := Job{0, 3, "hello", "T1"}
	world := Job{0, 5, "world", "T2"}
	go outputText(&hello)
	go outputText(&world)

}
