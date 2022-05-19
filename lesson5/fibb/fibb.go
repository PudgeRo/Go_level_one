package fibb

func Fibonacci(number int) int {
	if number == 0 || number == 1 {
		return number
	}
	return Fibonacci(number - 1) + Fibonacci(number - 2)
}

func FibonacciWithOptimization(number int, cache map[int]int) int {
	if val, ok := cache[number]; ok {
		return val
	}
	cache[number] = FibonacciWithOptimization(number - 1, cache) + FibonacciWithOptimization(number - 2, cache)
	return cache[number]
}
