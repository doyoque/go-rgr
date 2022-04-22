package main

import "fmt"

const messagePrefix = "Hello "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return messagePrefix + name + "!"
}

func main() {
	fmt.Println(Hello("Doyoque"))
}
