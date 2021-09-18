package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://www.prospecat.com",
		"http://golang.org",
		"http://stackoverflow.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	for {
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is maybe down!")
		c <- link
		return
	}
	fmt.Println(link, "is up and running!")
	c <- link
}
