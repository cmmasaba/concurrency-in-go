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

func fanIn(ch1, ch2 chan string) chan string {
	ch := make(chan string)

	go func ()  {
		for {
			select{
				case s := <-ch1: ch <- s
				case s := <-ch2: ch <- s
			}
		}
	} ()

	return ch

}

func main() {
	ch := fanIn(greet("Collins"), greet("Mmasaba"))

	for i := 0; i < 5; i++ {
		fmt.Println(<- ch)
	}
}