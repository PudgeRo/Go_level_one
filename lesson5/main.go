package main

import "fmt"

func main() {
	// The first solve
	fmt.Println(fibonacci1(5))

	// The second solve
	numbers := map[int]int{1:1, 2:1}
	fmt.Println(fibonacci2(5, numbers))
}

func fibonacci1(number int) int {
	if number == 0 || number == 1 {
		return number
	}
	return fibonacci1(number - 1) + fibonacci1(number - 2)
}

func fibonacci2(number int, numbers map[int]int) int {
	if number == 0 || number == 1 {
		return number
	}
	result := fibonacci2(number - 1, numbers) + fibonacci2(number - 2, numbers)
	numbers[number] = result
	return result
}
