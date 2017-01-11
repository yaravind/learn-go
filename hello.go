package main //package main

import (
	"fmt"

	"github.com/yaravind/example"
)

func main() {
	fmt.Println("Example - Hello World")
	fmt.Println("Hello")
	fmt.Println()

	fmt.Println("Example - Multiple value return from function")
	n, a := example.NameAndAge()
	fmt.Println(n, a)
}
