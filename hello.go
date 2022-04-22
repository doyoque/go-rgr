package main

import "fmt"

const messagePrefix = "Hello "
const messagePrefixSpain = "Hola "
const messagePrefixFrench = "Bonjour "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = messagePrefixFrench
	case "Spanish":
		prefix = messagePrefixSpain
	default:
		prefix = messagePrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Doyoque", ""))
}
