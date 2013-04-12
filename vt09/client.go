package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	server := []string{
	 	"http://localhost:8080",
	 	"http://localhost:8081",
	 	"http://localhost:8082",
	}
	for {
		before := time.Now()
		//res := Get(server[0])
		//res := Read(server[0], 0)
		res := MultiRead(server, time.Second)
		after := time.Now()
		fmt.Println("Response:", *res)
		fmt.Println("Time:", after.Sub(before))
		fmt.Println()
		time.Sleep(500 * time.Millisecond)
	}
}

type Response struct {
	Body       string
	StatusCode int
}

// Get makes an HTTP Get request and returns an abbreviated response.
// Status code 200 means that the request was successful.
// The function returns &Response{"", 0} if the request fails
// and it blocks forever if the server doesn't respond.
func Get(url string) *Response {
	res, err := http.Get(url)
	if err != nil {
		return &Response{}
	}
	// res.Body != nil when err == nil
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
	}
	return &Response{string(body), res.StatusCode}
}

// FIXME
// I've found two insidious bugs in this function; both of them are unlikely
// to show up in testing. Please fix them right away - and don't forget to
// write a doc comment this time.

// Bug 1: Both the main Go routine and the one calling Get() have read access 
// to res, which could result in a datarace. Could be fixed by changing it so 
// they do not access the same memory. 

// Read makes a HTTP Get request to the url and returns the response. Status
// code 200 means success. If the server does not answer within the suppied 
// timeout Read returns a Response stating timeout with status code 504. 
func Read(url string, timeout time.Duration) *Response {
	done := make(chan bool)
	res := new(Response)
	go func() {
		res = Get(url)
		done <- true
	}()
	select {
	case <-done:
	case <-time.After(timeout):
		return &Response{"Gateway timeout\n", 504}
	}
	return res
}

// MultiRead makes an HTTP Get request to each url and returns
// the response of the first server to answer with status code 200.
// If none of the servers answer before timeout, the response is
// 503 - Service unavailable.
	func MultiRead(urls []string, timeout time.Duration) (res *Response) {
	var responses = make([]*Response, len(urls))
	ch := make(chan int)
	for i, url := range urls {
		go func() {
			responses[i] = Get(url)
			ch <- i
		}()
	}
	select {
	case i := <-ch:
		res = responses[i]
	case <-time.After(timeout):
		res = &Response{"Service unavailable\n", 503}
	}
	return
}