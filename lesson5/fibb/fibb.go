package fibb

import "errors"

var ErrNonPositive = errors.New("number must be positive")

func Fibonacci(number int) (int, error) {
	if number < 0 {
		return 0, ErrNonPositive
	}
	if number == 0 || number == 1 {
		return number, nil
	}
	a, err := Fibonacci(number - 1)
	if err != nil {
		return 0, err
	}
	b, err := Fibonacci(number - 2)
	if err != nil {
		return 0, err
	}
	result := a + b
	return result, nil
}

func FibonacciWithOptimization(number int, cache map[int]int) (int, error) {
	if number < 0 {
		return 0, ErrNonPositive
	}
	if val, ok := cache[number]; ok {
		return val, nil
	}
	a, err := FibonacciWithOptimization(number-1, cache)
	if err != nil {
		return 0, err
	}
	b, err := FibonacciWithOptimization(number-2, cache)
	if err != nil {
		return 0, err
	}
	cache[number] = a + b
	return cache[number], nil
}
