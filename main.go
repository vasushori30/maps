package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"while": "#ffffff",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	// Here color is key and hex is value
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
