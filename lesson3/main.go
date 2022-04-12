package main

import (
	"bufio"
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
			actions := strings.Split(expression, " ")
			if len(actions) > 3 || len(actions) == 0 || len(actions) == 2 {
				fmt.Println("Incorrectly entered data")
			} else if len(actions) == 1 {
				if strings.HasSuffix(actions[0], "!")  {
					actions[0] = strings.Trim(actions[0], "!")
					number1, errFirstNum := strconv.Atoi(actions[0])
					if errFirstNum == nil {
						fmt.Println(expression, "=", factorial(number1))
						continue
					} else {
						fmt.Println("Invalid types of number")
						continue
					}
				} else {
					fmt.Println("Incorrectly entered data")
				}
			} else {			
				number1, errFirstNum := strconv.ParseFloat(actions[0], 64)
				number2, errSecondNum := strconv.ParseFloat(actions[2], 64)
				if errFirstNum == nil && errSecondNum == nil{
					if (actions[1] == "+") || (actions[1] == "-") || (actions[1] == "*") || (actions[1] == "^") {
						fmt.Printf("%s = %.2f\n", expression, calculator(number1, number2, actions[1]))
					} else if actions[1] == "/" {
						if number2 != 0 {
							fmt.Printf("%s = %.2f\n", expression, calculator(number1, number2, actions[1]))
						} else {
							fmt.Println("Division by zero")
						}
					} else {
						fmt.Println("Invalid types of operation")
					}
				} else {
					fmt.Println("Invalid types of number(-s)")
				}
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

func calculator(firstNum, secondNum float64, operation string) float64 {
	switch operation {
	case "+": 
		return firstNum + secondNum
	case "-": 
		return firstNum - secondNum
	case "*": 
		return firstNum * secondNum
	case "/": 
		return firstNum / secondNum
	case "^": 
		return math.Pow(firstNum, secondNum)
	default:
		return 0
	}
}

func factorial(number int) int {
	var result int = 1
	for i := 1; i <= number; i++ {
		result *= i
	}
	return result
}