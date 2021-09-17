package main

import "fmt"

type bot interface{
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb:=englishBot{}
	es:=spanishBot{}
	printGreeting(eb)
	printGreeting(es)
}

func (englishBot) getGreeting() string {
	return "Hello"
}
func (spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}