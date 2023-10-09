package main

import "fmt"

func main() {
	colors := make(map[int]string)
	colors[5] = "#ff0000"
	delete(colors, 5)
	fmt.Println(colors[5])
}
