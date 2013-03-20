package main

import (
	"fmt"
	"time"
)

func Remind(text string, paus time.Duration) {
	for {
		fmt.Println("Klockan är " + time.Now().Format("15:04:05") + ": " + text)
		time.Sleep(paus)
	}
}     

func main() {
	Remind("Hej!", time.Second)
}
     