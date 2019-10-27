package main

import "fmt"

var languageMap = map[string]string{
	"Spanish": "Hola, ",
	"English": "Hello, ",
	"French":  "Bonjour, ",
}

func Hello(name, language string) string {

	if greeting, ok := languageMap[language]; ok {
		return greeting + name
	}

	language = "English"
	var greeting = languageMap[language]

	if name != "" {
		return greeting + name
	}

	return greeting + "world"
}

func main() {
	fmt.Println(Hello("world", ""))
}
