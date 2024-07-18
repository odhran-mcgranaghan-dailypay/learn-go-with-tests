package context

import (
	"context"
	"fmt"
)

// Original version
func FOriginal(ctx context.Context) {
	context.WithValue(ctx, "foo", -6)
}

func Original() {
	ctx := context.TODO()
	FOriginal(ctx)
	fmt.Println("Original:", ctx.Value("foo")) // Output: <nil>
}

// Corrected version
func FCorrected(ctx context.Context) context.Context {
	return context.WithValue(ctx, "foo", -6)
}

func Corrected() {
	ctx := context.TODO()
	ctx = FCorrected(ctx)                       // Update ctx with the new context returned by f
	fmt.Println("Corrected:", ctx.Value("foo")) // Output: -6
}

func run() {
	Original()
	Corrected()
}
