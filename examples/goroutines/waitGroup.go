package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
	threadName string
}

func outputText(j *Job, goGroup *sync.WaitGroup) {
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text + " : " + j.threadName)
		j.i++
	}
	goGroup.Done()
}
func main() {
	goGroup := new(sync.WaitGroup)
	hello := Job{0, 3, "hello", "T1"}
	world := Job{0, 5, "world", "T2"}
	go outputText(&hello, goGroup)
	go outputText(&world, goGroup)
	goGroup.Add(2)
	goGroup.Wait()


}
