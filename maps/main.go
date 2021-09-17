package main

import "fmt"

func main() {
	// var colors map[string]string
	colors := make(map[string]string)
	colors["white"] = "#fff"
	delete(colors, "white")
	// colors := map[string]string{
	// 	"yellow": "#FFFF00",
	// 	"white":  "#FFFFFF",
	// }
	fmt.Println(colors)
}
