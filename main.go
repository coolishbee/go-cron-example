package main

import (
	"fmt"
	"runtime"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("iam running task.")
}

func main() {
	currentCPU := runtime.NumCPU()
	fmt.Println("count : ", currentCPU)
	//runtime.GOMAXPROCS(currentCPU)

	gocron.Every(1).Second().Do(task)
	<-gocron.Start()
}
