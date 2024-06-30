package main

import (
	"bytes"
	"testing"
)

// print 3, 321, and the spaced out by second

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := "3\n2\n1\nGo!\n"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	// dependency injection
	// test the mocked out Sleep function using spySleeper
	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}

}
