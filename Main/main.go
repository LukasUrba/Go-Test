package main

import (
	text "fmt"

	helloWorld "github.com/LukasUrba/Go/helloWorld"
	values "github.com/LukasUrba/Go/values"
)

func main() {
	number := values.Numbers()
	helloWorld.Run()
	text.Println(number)
	// test
}
