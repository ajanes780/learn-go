package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	c := make(chan string)
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.stackoverflow.com",
		"https://www.golang.org",
		"https://www.youtube.com",
	}

	for _, link := range links {
		// Go basically creates a new thread /sub routine for each function call.
		go checkLink(link, c)
	}

	// This syntax creates a loop for the range of the channel c.
	for l := range c {

		// This is a function literal. It is a function that is not declared.
		// this is the same as and anonymous function in javascript.
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
