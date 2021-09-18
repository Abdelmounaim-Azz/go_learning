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
	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is maybe down!")
		return
	}
	fmt.Println(link, "is up and running!")
}
