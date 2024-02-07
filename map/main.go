package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(myMap map[string]string) {
	for i, color := range myMap {
		fmt.Println("Color", i, "is", color)
	}
}
