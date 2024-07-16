package reflection

import (
	"fmt"
	"reflect"
	"testing"
)

// https://go.dev/blog/laws-of-reflection

// TestSettability demonstrates the third law of reflection: to modify a reflection object, the value must be settable.

// It’s the property that a reflection object can modify the actual storage that was used to create the reflection object. 
// Settability is determined by whether the reflection object holds the original item. When we say


func TestSettability(t *testing.T) {
	// This test case will panic because the reflection object is not settable.
	t.Run("non-settable value", func(t *testing.T) {
		var x float64 = 3.4
		v := reflect.ValueOf(x)

		// Check and print the settability of v
		settable := v.CanSet()
		fmt.Println("settability of v:", settable)

		if settable {
			t.Errorf("expected settability of v to be false, got true")
		}

		// Uncommenting the following line will cause a panic
		// v.SetFloat(7.1)
	})

	// This test case shows that to modify a reflection object, we must pass a pointer to the value.
	t.Run("settable value", func(t *testing.T) {
		var x float64 = 3.4
		p := reflect.ValueOf(&x) // Pass the address of x

		// Print the type of p
		fmt.Println("type of p:", p.Type())

		// Check and print the settability of p
		settableP := p.CanSet()
		fmt.Println("settability of p:", settableP)

		if settableP {
			t.Errorf("expected settability of p to be false, got true")
		}

		// The reflection object p isn’t settable, but it’s not p we want to set, it’s (in effect) *p. 
		// To get to what p points to, we call the Elem method of Value, which indirects through the pointer, 
		// and save the result in a reflection Value called v:

		// Use Elem to get the reflection object that p points to
		v := p.Elem()

		// Check and print the settability of v
		settableV := v.CanSet()
		fmt.Println("settability of v:", settableV)

		if !settableV {
			t.Errorf("expected settability of v to be true, got false")
		}

		// Now v is settable and we can modify the value of x through v
		v.SetFloat(7.1)

		// Verify that the value of x has been updated
		if x != 7.1 {
			t.Errorf("expected x to be 7.1, got %v", x)
		}

		// Print the final values
		fmt.Println("v.Interface():", v.Interface())
		fmt.Println("x:", x)
	})
}
