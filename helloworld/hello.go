package helloworld

import "fmt"

// Microlesson #1
// group related constants for readability
const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

// Microlesson #2
// using map instead of switch statement
// if you had many switch statement cases, I think a map is bit more readable
// I didn't benchmark this, but I my guess is it would be faster than a switch statement O(n) vs O(1) lookup
var greetings = map[string]string{
	spanish: spanishHelloPrefix,
	french:  frenchHelloPrefix,
}

func greetingPrefix(language string) (prefix string) {
	exists := false
	if prefix, exists = greetings[language]; !exists {
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
