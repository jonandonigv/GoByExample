package main

import (
	"fmt"
	"time"
)

func main() {
	// create a timer which will star 2s after we start the program
	timer1 := time.NewTimer(2 * time.Second)

	// blocks on the timer's channel until it sends a value
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// create a timer which will start after one second
	timer2 := time.NewTimer(time.Second)
	// create a go routine that blocks the timer channel until a val is sended
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// we stop the timer before it fires
	stop2 := timer2.Stop()
	// The timer was stoped before firing so we println
	if stop2 {
		fmt.Println("Timer 2 stoppped")
	}

	time.Sleep(2 * time.Second)
}
