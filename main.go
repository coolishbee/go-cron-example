package main

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("iam running task.")
}

func main() {
	gocron.Every(1).Second().Do(task)
	<-gocron.Start()
}
