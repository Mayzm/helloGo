package main

import (
	"fmt"
	"runtime"
	"os"
)

func main() {
	a, b := test(2)
	fmt.Printf("a = %d, b = %d\n", a, b)
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}

func test(a) (int, int) {
	key, value := 1, a
	return key, value
}
