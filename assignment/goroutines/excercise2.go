package main

import (
	"fmt"
)

// 2. Concurrent Sum Calculator
// Problem:

// Write a program that:
// Splits a slice of numbers into two halves
// Calculates the sum of each half in separate goroutines
// Combines the results and prints the total sum

func sum(subArr []int, ch chan int) {
	sum := 0
	for _, v := range subArr {
		sum += v
	}
	ch <- sum
}

func runConcurrentSumCalculator() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	numArr := []int{5, 7, 0, 2, 3, 4, 5, 6}

	mid := len(numArr) / 2
	go sum(numArr[0:mid], ch1)
	go sum(numArr[mid:], ch2)

	total := <-ch1 + <-ch2

	fmt.Println("main ends: total:", total)
}
