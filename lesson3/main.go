package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Simple number from 0 to N
	// simpleNumbers(19)

	// Calculator
	stop := true
	for stop {
		fmt.Println("Enter the expression in this format: 1 + 1 (number1 operation number2) ")
		fmt.Println("Avilable operation '+', '-', '*', '/', '^' (the number1 is raised to the power of the number2), '!' (input formal: 1!)")
		fmt.Println("To stop calculate enter 'stop'")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			expression := scanner.Text()
			if expression == "stop" {
				stop = false
				break
			}
			actions := strings.Split(expression, " ") // разбиваем по пробелам введенный текст на символы и вносим их в слайс
			if len(actions) > 3 || len(actions) == 0 || len(actions) == 2 {
				fmt.Println("incorrect expression")
			} else if len(actions) == 1 {
				if strings.HasSuffix(actions[0], "!")  {
					actions[0] = strings.Trim(actions[0], "!")
					number1, err := strconv.Atoi(actions[0])
					if err == nil {
						fmt.Printf("Result: %v\n", factorial(number1))
						continue
					} else {
						fmt.Println("the number must be an integer")
						continue
					}
				} else {
					fmt.Println("incorrect expression")
				}
			} else {		
				number1, err := strconv.ParseFloat(actions[0], 64)
				if err != nil {
					fmt.Println("incorrect data")
					continue 
				}
				number2, err := strconv.ParseFloat(actions[2], 64)
				if err != nil {
					fmt.Println("incorrect data")
					continue
				}

				op, err := operation(actions[1])
				if err != nil {
					fmt.Printf("incorrect data: %s\n", err)
					continue 
				}
				res, err := calculator(number1, number2, op)
				if err != nil {
					fmt.Println(err)
					continue 
				}
				fmt.Printf("Result: %.2f\n", res)
			}
		}	
	}
}
// func simpleNumbers (N int) {
// 	var isSimple bool
// 	for i := 1; i <= N; i++ {
// 		isSimple = true
// 		for j := 2; j < i; j++ {
// 			if i % j == 0 {
// 				isSimple = false
// 				break
// 			}
// 		}
// 		if isSimple {
// 			fmt.Printf("%d- простое число\n", i)
// 		}
// 	}
// }

func calculator(firstNum, secondNum float64, operation string) (float64, error) {
	var result float64
	switch operation {
	case "+": 
		result = firstNum + secondNum
	case "-": 
		result = firstNum - secondNum
	case "*": 
		result = firstNum * secondNum
	case "/": 
		divideResult, err := divide(firstNum, secondNum); 
		if err != nil {
			return 0, fmt.Errorf("could not divide: %w", err)
		}
		result = divideResult
	case "^": 
		result = math.Pow(firstNum, secondNum)
	}
	return result, nil
}

func divide(firstNum, secondNum float64) (float64, error) {
	if secondNum == 0 {
		return 0, errors.New("division by zero")
	}
	return firstNum / secondNum, nil
}

func factorial(number int) int {
	var result int = 1
	for i := 1; i <= number; i++ {
		result *= i
	}
	return result
}

func operation(op string) (string, error) {
	if op == "+" || op == "-" || op == "*" || op == "/" {
		return op, nil
	}
	return "", errors.New("incorrect operation")
}