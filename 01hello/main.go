package main

import "fmt"

const (
 hindi = "Hindi"
 spanish = "Spanish"
 englishHelloPrefix = "Hello "
 hindiHelloPrefix = "Namaste "
 spanishHelloPrefix = "Hola "
)

func Hello(name string, lang string) string{
	if name == "" {
		name = "World"
	}
	
	return greetingPrefix(lang) + name + "."
}

func greetingPrefix(lang string) (prefix string){
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case hindi:
		prefix = hindiHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main(){
	fmt.Println(Hello("Armaan", "Spanish"))
}