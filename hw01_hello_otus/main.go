package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	newStr := "Hello, OTUS!"
	fmt.Printf("%s\n", reverse.String(newStr))
}
