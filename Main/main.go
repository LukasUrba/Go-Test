package main

import (
	text "fmt"

	helloWorld "example.com/mainFile/helloWorld"
	values "example.com/mainFile/values"
)

func main() {
	number := values.Numbers()
	helloWorld.Run()
	text.Println(number)
	// test
}
