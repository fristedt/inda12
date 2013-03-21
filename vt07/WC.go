package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    mappy := make(map[string]int)
    words := strings.Fields(s)
    for _, word := range words {
    	mappy[word] += 1
    }
    return mappy
}

func main() {
    wc.Test(WordCount)
}

