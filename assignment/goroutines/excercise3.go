package main

import (
	"fmt"
	"time"
)

// 3. Ping-Pong with Channels
// Problem:

// Create two goroutines:

// "Ping" sends "ping" to a channel
// "Pong" receives "ping" and responds with "pong"
// Repeat 3 times before exiting

func ping(ch1 chan string) {
	fmt.Println("Ping")
	ch1 <- "Ping"
}

func pong(ch1 chan string) {
	<-ch1
	fmt.Println("Pong")
	ch1 <- "Pong"
}

func runPingPong() {
	ch1 := make(chan string)

	for i := 0; i < 3; i++ {
		go ping(ch1)
		go pong(ch1)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("main ends")
}
