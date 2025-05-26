package main

import "fmt"

func greet(s string) chan string{
	ch := make(chan string)

	go func() {for i:=0; ;i++{
			ch <- fmt.Sprintf("%d. hello, %s", i, s)
		}
	} ()
	return ch
}

func main() {
	collins := greet("Collins")

	for i := 0; i < 5; i++ {
		fmt.Println(<- collins)
	}
}