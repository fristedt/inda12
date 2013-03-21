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
	go Remind("Dags att äta", time.Hour*3)
	go Remind("Dags att arbeta", time.Hour*8)
	Remind("Dags att sova", time.Hour*24)
}
