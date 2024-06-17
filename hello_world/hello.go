package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "spanish"))
}

// Notes:
// In our function signature we have made a named return value (prefix string).
// This will create a variable called prefix in your function.
// It will be assigned the "zero" value. This depends on the type, for example ints are 0 and for strings it is "".
// You can return whatever it's set to by just calling return rather than return prefix.
// This will display in the Go Doc for your function so it can make the intent of your code clearer.
// default in the switch case will be branched to if none of the other case statements match.
// The function name starts with a lowercase letter. In Go, public functions start with a capital letter, and private ones start with a lowercase letter. We don't want the internals of our algorithm exposed to the world, so we made this function private.
// Also, we can group constants in a block instead of declaring them on their own line. For readability, it's a good idea to use a line between sets of related constants.
