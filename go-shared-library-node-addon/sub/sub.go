package main

import (
	"C"
)

//export sub
func sub(a, b int) int {
	return a - b
}

func main() {
}
