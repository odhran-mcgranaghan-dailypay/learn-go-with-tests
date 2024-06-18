package iteration

func Sum(numbers []int) int {
	var result int
	// size is encoded in its type. [4]int != [5]int, it won't compile.

	// using range to interate over the array, tuple assignment
	// using range here returns two values, the index and the value
	// using _ (blank identifier) ignores the index, we only care about the value so add to the result
	for _, number := range numbers {
		result += number
	}
	return result
}

// Slices - An array is a fixed-size collection of elements of a single type,
// 			while a slice is a dynamically-sized, flexible view into the elements of an array.
// mySlice := []int{1,2,3} rather than myArray := [3]int{1,2,3}
