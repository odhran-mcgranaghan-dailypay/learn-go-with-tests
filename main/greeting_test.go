package main

import (
	"bytes"
	"io"
	"os"
	"testing"
	"net"
)

func TestGreet(t *testing.T) {
	t.Run("Test Case 1: Writing to a Buffer", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "chris")

		got := buffer.String()
		want := "Hello, chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Test Case 2: File writer", func(t *testing.T) {
		// create a temp file
		tempFile, err := os.CreateTemp("", "greetings.txt")
		if err != nil {
			t.Fatal("could not create temp file")
		}
		// defer here ensures that the file is closed and removed after the function completes
		defer os.Remove(tempFile.Name())

		Greet(tempFile, "Scooby Doo")

		// read the file, setting pointer to start of file
		// io.SeekStart is a constant that represents the starting point of the file
		tempFile.Seek(0, io.SeekStart)
		content, err := io.ReadAll(tempFile)
		if err != nil {
			t.Fatalf("could not read temp file")
		}

		// convert byte slice to string
		got := string(content)
		want := "Hello, Scooby Doo"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("Test Case 3: Writing to a Network", func(t *testing.T) {
		client, server := net.Pipe()
		defer client.Close()
		defer server.Close()

		// Write to the network using the server end
		go Greet(server, "network")

		// Read from the network using the client end
		buf := new(bytes.Buffer)
		io.Copy(buf, client)

		got := buf.String()
		want := "Hello, network"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
