package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}
	c <- sum   // Send to channel
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)

	go sum(s[len(s)/2:], ch)
	go sum(s[:len(s)/2], ch)

	x, y := <-ch, <-ch   // Receive from channel
	fmt.Printf("x: %v y: %v: %v\n", x, y, x+y)
} 