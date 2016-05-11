package main

import (
	// "fmt"
	"time"
)

func readqueue(queueUrl string) {
	ticker := time.NewTicker(time.Second * 2)
	for range ticker.C {
		// fmt.Println("Tick 1 at", time.Now())
		getMessages(queueUrl)
		// fmt.Println("Tick 2 at", time.Now())
		// fmt.Println("Tick T at", t)
	}

}
