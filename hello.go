// Concepts of TDD, testing, const, switch, if else

package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, langauge string) string{

	if name == "" {
		name = "World"
	}

	return greetingPrefix(langauge) + name
}

func greetingPrefix(language string) string {
	switch language {
		case spanish:
			return spanishHelloPrefix
		case french:
			return frenchHelloPrefix
		default:
			return englishHelloPrefix
	}
}

func main() {
	fmt.Print(Hello("Chris", ""))
}