package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
// You need to send the string to the channel in another thread.
func main() {
	ch := make(chan string)
	go func() {
		ch <- "Hello world!"
	}()
	fmt.Println(<-ch)
}
