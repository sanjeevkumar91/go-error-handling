package main

import (
	"fmt"
	"time"
)

// Create a program that launches 5 goroutines. Each goroutine should:
// Take a name as input (e.g., "Alice", "Bob")
// Sleep for a random time (1-3 seconds)
// Print "Hello, [name]!"

func printName(name string, timeout int) {
	time.Sleep(time.Duration(timeout) * time.Second)
	fmt.Println(name)
}

func printNames() {
	go printName("Alice", 2)
	go printName("Bob", 3)
	go printName("John", 1)
	go printName("Ram", 2)
	go printName("Rajesh", 1)
	fmt.Println("main ends")
	time.Sleep(4 * time.Second)
}
