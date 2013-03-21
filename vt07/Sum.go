package main

import "fmt"

// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan<- int) {
	sum := 0
	for _, value := range a {
		sum += value
	}
	res <- sum
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}

	n := len(a)
	ch := make(chan int)
	go Add(a[:n/2], ch)
	go Add(a[n/2:], ch)

	// TODO: Get the subtotals from the channel and print their sum.
	sum1 := <-ch
	sum2 := <-ch
	fmt.Printf("%d %d %d \n", sum1, sum2, sum1 + sum2)
}
