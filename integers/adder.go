package integers

func Add(x, y int) int {
	return x + y
}

// Created a variadic function that takes multiple integers and returns the sum of all the integers
func AddMultiple(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
