package main

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "Spanish" {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}
