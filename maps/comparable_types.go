package maps

import (
	"fmt"
)

// A map is declared using the map keyword followed by the key and value types
// e.g. map[string]int
// e.g. map["UK"] = 44, map["Finland"] = 358
// The zero value of a map is nil
// The key of a map must be a type that is comparable
// To compare two values they must be Assignable to each other: https://go.dev/ref/spec#Assignability
// Below are examples of all types that are comparable in Go: https://go.dev/ref/spec#Comparison_operators
// The value may be of any type, including another map
func main() {
	// Boolean comparison
	b1 := true
	b2 := false
	fmt.Println("b1 == b2:", b1 == b2) // false
	fmt.Println("b1 != b2:", b1 != b2) // true

	// Integer comparison
	var x, y int = 5, 10
	fmt.Println("x == y:", x == y) // false
	fmt.Println("x < y:", x < y)   // true
	fmt.Println("x <= y:", x <= y) // true
	fmt.Println("x > y:", x > y)   // false
	fmt.Println("x >= y:", x >= y) // false

	// Floating-point comparison
	var f1, f2 float64 = 3.14, 6.28
	fmt.Println("f1 == f2:", f1 == f2) // false
	fmt.Println("f1 < f2:", f1 < f2)   // true

	// Complex number comparison
	var c1, c2 complex64 = complex(1, 2), complex(1, 2)
	fmt.Println("c1 == c2:", c1 == c2) // true
	// c3 := complex(2, 3)
	// fmt.Println("c1 == c3:", c1 == c3) // false

	// String comparison
	s1 := "apple"
	s2 := "banana"
	fmt.Println("s1 == s2:", s1 == s2) // false
	fmt.Println("s1 < s2:", s1 < s2)   // true

	// Pointer comparison
	var p1, p2 *int
	fmt.Println("p1 == p2:", p1 == p2) // true
	p1 = &x
	p2 = &y
	fmt.Println("p1 == p2:", p1 == p2) // false
	p2 = &x
	fmt.Println("p1 == p2:", p1 == p2) // true

	// Channel comparison
	ch1 := make(chan int)
	ch2 := make(chan int)
	fmt.Println("ch1 == ch2:", ch1 == ch2) // false
	ch3 := ch1
	fmt.Println("ch1 == ch3:", ch1 == ch3) // true

	// Interface comparison
	var i1, i2 interface{}
	i1 = 42
	i2 = 42
	fmt.Println("i1 == i2:", i1 == i2) // true
	i2 = "hello"
	fmt.Println("i1 == i2:", i1 == i2) // false

	// Struct comparison
	type Point struct {
		X, Y int
	}
	point1 := Point{X: 1, Y: 2}
	point2 := Point{X: 1, Y: 2}
	point3 := Point{X: 2, Y: 3}
	fmt.Println("point1 == point2:", point1 == point2) // true
	fmt.Println("point1 == point3:", point1 == point3) // false

	// Array comparison
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{4, 5, 6}
	fmt.Println("arr1 == arr2:", arr1 == arr2) // true
	fmt.Println("arr1 == arr3:", arr1 == arr3) // false

	// Special cases
	// Slices, maps, and functions are not comparable except with nil
	var s1Slice []int
	fmt.Println("s1Slice == nil:", s1Slice == nil) // true

	var m1Map map[string]int
	fmt.Println("m1Map == nil:", m1Map == nil) // true

	var f1Func func()
	fmt.Println("f1Func == nil:", f1Func == nil) // true

	// Untyped boolean constant
	const c = 3 < 4
	fmt.Println("const c = 3 < 4:", c) // true

	// Comparison result in different types
	type MyBool bool
	var b3 = x == y
	var b4 bool = x == y
	var b5 MyBool = x == y
	fmt.Println("b3 (bool):", b3)   // false
	fmt.Println("b4 (bool):", b4)   // false
	fmt.Println("b5 (MyBool):", b5) // false
}
