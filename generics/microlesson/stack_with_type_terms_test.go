package microlesson

import (
	"testing"
)

type MySpecialString string

func TestStack(t *testing.T) {

	t.Run("Type Term - string", func(t *testing.T) {
		stack := new(StackV2[string])
		stack.Push("hello")
		val, _ := stack.Pop()
		AssertEqual(t, val, "hello")

	})

	t.Run("Type Term - int", func(t *testing.T) {
		stack := new(StackV2[int])
		stack.Push(10)
		val, _ := stack.Pop()
		AssertEqual(t, val, 10)

	})

	t.Run("Type Term - underlying Type test", func(t *testing.T) {
		stack := new(StackV2[MySpecialString])
		stack.Push("hello")
		val, _ := stack.Pop()
		AssertEqual(t, val, "hello")

	})

	// interfaces with type elements are valid only as type constraints.
	// It is a compile-time error to use them as the type for a variable, field, return value, or parameter.

	t.Run("Type Term - illegal usage of type terms", func(t *testing.T) {
		//
		var myVar AcceptableTypeTerms
		myVar = 42
		myVar = "hello"
	})

	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(StackV2[int])

		// check stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())

		// add a thing, then check it's not empty
		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty())

		// add another thing, pop it back again
		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 456)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, myStackOfInts.IsEmpty())
	})

	t.Run("string stack", func(t *testing.T) {
		stackOfStrings := new(StackV2[string])

		AssertTrue(t, stackOfStrings.IsEmpty())
		stackOfStrings.Push("odhran")
		AssertFalse(t, stackOfStrings.IsEmpty())

		stackOfStrings.Push("dog")
		value, _ := stackOfStrings.Pop()
		AssertEqual(t, value, "dog")
		value, _ = stackOfStrings.Pop()
		AssertTrue(t, stackOfStrings.IsEmpty())

	})
}

// Assert helper functions
func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
