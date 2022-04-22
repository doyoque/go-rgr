package main

import "fmt"

const messagePrefix = "Hello "
const messagePrefixSpain = "Hola "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "Spanish" {
		return messagePrefixSpain + name + "!"
	}

	return messagePrefix + name + "!"
}

func main() {
	fmt.Println(Hello("Doyoque", ""))
}
