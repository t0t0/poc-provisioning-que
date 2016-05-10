package main

import (
	"fmt"
	"time"
)

func readqueue() {
	ticker := time.NewTicker(time.Second * 2)
	for t := range ticker.C {
		fmt.Println("Tick 1 at", time.Now())
		getMessages()
		fmt.Println("Tick 2 at", time.Now())
		fmt.Println("Tick T at", t)
	}

}
