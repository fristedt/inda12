package main

import "fmt"

// This program should go to 11, but sometimes it only prints 1 to 10.
/* The bug was that sometimes the main method exited before the Print method got a chance to print. In this version the for loop that prints is in the main function, so the program won't exit before all values have been printed. */
func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 11; i++ {
			ch <- i
		}
		close(ch)
	}()
	for n := range ch { // reads from channel until it's closed
		fmt.Println(n)
	}
}
