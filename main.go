package main

import (
	"fmt"

	go_say_hello "github.com/mhaatha/go-say-hello"
)

func main() {
	message := go_say_hello.SayHello("User")
	fmt.Println(message)
}
