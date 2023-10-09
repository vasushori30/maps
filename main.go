package main

import "fmt"

func main() {
	colors := make(map[int]string)
	colors[5] = "#ff0000"
	fmt.Println(colors[5])
}
