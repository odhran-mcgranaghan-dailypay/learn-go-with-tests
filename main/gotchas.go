package main

import (
	"context"
	"fmt"
)

func f(ctx context.Context) (int) {
	context.WithValue(ctx, "foo", -6)
	return 0
}

func main() {
	ctx := context.TODO()
	f(ctx)
	fmt.Println(ctx.Value("foo"))
}
