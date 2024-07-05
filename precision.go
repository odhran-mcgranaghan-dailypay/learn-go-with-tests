package main

import (
	"fmt"
)

// link to docs https://pkg.go.dev/fmt

func Run() {
	value := 123.456789

	// .g - It prints the number with the shortest representation that preserves the specified precision.

	fmt.Printf("%.g\n", value)  // 1e+02
	fmt.Printf("%.2g\n", value) // 1.2e+02
	fmt.Printf("%.5g\n", value) // 123.46
	fmt.Printf("%.8g\n", value) // 123.45679

	// .f - It prints the number with the specified precision, a fixed number of digits after the decimal point.

	fmt.Printf("%.f\n", value)  // 123.45679
	fmt.Printf("%.2f\n", value) // 123.46
	fmt.Printf("%.5f\n", value) // 123.45679

	// Use %f when you want to control the number of digits after the decimal point in decimal notation.
	// Use %g when you want a more flexible representation that can switch between decimal and scientific notation, aiming to use the shortest form while preserving the specified precision.
}
