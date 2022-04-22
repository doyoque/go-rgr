package main

import "fmt"

const messagePrefix = "Hello "
const messagePrefixSpain = "Hola "
const messagePrefixFrench = "Bonjour "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := messagePrefix

	switch language {
	case "French":
		prefix = messagePrefixFrench
	case "Spanish":
		prefix = messagePrefixSpain
	}

	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("Doyoque", ""))
}
