package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	// approach 2
	fmt.Fprintf(writer, "Hello, %s", name)

}

func MyGreetingHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// Use cases and test cases for these

// func main() {
// 	// Write output to standard output
// 	Greet(os.Stdout, "John")

// 	// Write output to a network
// 	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetingHandler)))

// 	// Write output to a file
// 	filename := "greetings.txt"
// 	f, _ := os.Create(filename)
// 	// returns a *File, File implements the Writer interface, we can use the Greet function to write to the file
// 	Greet(f, "Scooby Doo")

// }

// Extensions: What other important Types in go implement the io.Writer interface?

// file writing - os.File

// extensions
// logging, file output, network output, standard output, buffer

// io.Writer general purpose interface for writing data, files, stdout, network, etc.
// Satisfy interface by implementing Write method
