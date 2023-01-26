package main

import (
	helloWorld "example.com/mainFile/helloWorld"
	values "example.com/mainFile/values"
	text "fmt"
)

func main() {
	number := values.Numbers()
	helloWorld.Run()
	text.Println(number)
}
