package main

import "fmt"

const messagePrefix = "Hello "
const messagePrefixSpain = "Hola "
const messagePrefixFrench = "Bonjour "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "Spanish" {
		return messagePrefixSpain + name + "!"
	}

	if language == "French" {
		return messagePrefixFrench + name + "!"
	}

	return messagePrefix + name + "!"
}

func main() {
	fmt.Println(Hello("Doyoque", ""))
}
