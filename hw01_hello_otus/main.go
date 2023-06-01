package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const hello string = "Hello, OTUS!"

func main() {
	fmt.Print(stringutil.Reverse(hello))
}
