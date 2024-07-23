package generics

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

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
		stackOfStrings := new(Stack[string])

		AssertTrue(t, stackOfStrings.IsEmpty())
		stackOfStrings.Push("odhran")
		AssertFalse(t, stackOfStrings.IsEmpty())

		stackOfStrings.Push("banbha")
		value, _ := stackOfStrings.Pop()
		AssertEqual(t, value, "banbha")
		value, _ = stackOfStrings.Pop()
		AssertTrue(t, stackOfStrings.IsEmpty())

	})
	t.Run("interface stack dx is horrid", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()

		// get our ints from out interface{}
		reallyFirstNum, ok := firstNum.(int)
		AssertTrue(t, ok) // need to check we definitely got an int out of the interface{}

		reallySecondNum, ok := secondNum.(int)
		AssertTrue(t, ok) // and again!

		AssertEqual(t, reallyFirstNum+reallySecondNum, 3)
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
