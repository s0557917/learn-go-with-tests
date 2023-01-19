package main

import "fmt"

const englishHelloPrefix = "Hello "
const englishHello = "Hello World!"

func Hello(name string) string {
	if name == "" {
		return englishHello
	} else {
		return englishHelloPrefix + name + "!"
	}

}

func main() {
	fmt.Println(Hello("Phillip"))
}
