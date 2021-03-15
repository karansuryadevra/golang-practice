package main

import "fmt"

const englishHelloPrefix = "Hello, "

func HelloWorld(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(HelloWorld("Karan"))
}
