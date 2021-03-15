package main

import "fmt"

const EnglishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "
const FrenchHelloPrefix = "Bonjour, "
const ItalianHelloPrefix = "Buongiorno, "
const Spanish = "Spanish"
const French = "French"
const Italian = "Italian"

func HelloWorld(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case French:
		prefix = FrenchHelloPrefix
	case Spanish:
		prefix = SpanishHelloPrefix
	case Italian:
		prefix = ItalianHelloPrefix
	default:
		prefix = EnglishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(HelloWorld("Karan", ""))
}
