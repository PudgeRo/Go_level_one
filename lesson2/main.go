package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// The area of the rectangle
	// var length, width float64
	// fmt.Print("Enter the length: ")
	// fmt.Scan(&length)
	// fmt.Print("Enter the width: ")
	// fmt.Scan(&width)
	// fmt.Printf("The area of the rectangle is %.2f\n\n", rectangle(length, width))

	// Diameter and circumference 
	// var circleArea float64
	// fmt.Print("Enter the area of a circle: ")
	// fmt.Scan(&circleArea)
	// diameter, circumference := circle(circleArea)
	// fmt.Printf("Diameter is %.2f, circumference is %.2f\n\n", diameter, circumference)

	// Decomposition of number
	// var number int
	// fmt.Print("Enter a three-digit number: ")
	// fmt.Scan(&number)
	// units, tens, hundredths := threeDigitNum(number)
	// fmt.Printf("Units: %d\nTens: %d\nHundredths: %d\n\n", units, tens, hundredths)

	// Find simple number from 0 to N
	// var N int
	// fmt.Print("Enter N: ")
	// fmt.Scan(&N)
	// simpleDigits(N)

	//Calculator 
	var number1, number2 string
	var op string
	fmt.Println("Enter 'stop' instead of first number for stopping program")
	for {
		fmt.Print("Enter first number: ")
		fmt.Scan(&number1)
		if number1 == "stop" {
			break
		}
		number1, err := strconv.ParseFloat(number1, 64)
		if err != nil {
			fmt.Println("incorrect data")
			continue
		}

		fmt.Print("Enter second number: ")
		fmt.Scan(&number2)
		number2, err := strconv.ParseFloat(number2, 64)
		if err != nil {
			fmt.Println("incorrect data")
			continue
		}

		fmt.Print("Enter operation: ")
		fmt.Scan(&op)		
		op, err := operation(op)
		if err != nil {
			fmt.Printf("incorrect data: %s\n", err)
			continue
		}

		res, err := calculator(number1, number2, op)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Result is %.2f\n", res)
	}
}

func operation(op string) (string, error) {
	if op == "+" || op == "-" || op == "*" || op == "/" {
		return op, nil
	}
	return "", errors.New("incorrect opearation")
}

func calculator(number1, number2 float64, operation string) (float64, error) {
	var result float64
	switch operation {
	case "+":
		result = number1 + number2
	case "-":
		result = number1 - number2
	case "*":
		result = number1 * number2
	case "/":
		divideResult, err := divide(number1, number2); 
		if err != nil {
			return 0, fmt.Errorf("could not divide: %w", err)
		}
		result = divideResult
	}
	return result, nil
}

func divide(number1, number2 float64) (float64, error) {
	if number2 == 0 {
		return 0, errors.New("division by zero")
	}
	return number1 / number2, nil
}

// func rectangle(length, width float64) float64{
// 	rectangleArea := length * width
// 	return rectangleArea
// }

// func circle(circleArea float64) (float64, float64){
// 	pi := math.Pi
// 	radius := math.Sqrt(circleArea / pi)
// 	circumference := 2 * pi * radius
// 	diameter := 2 * radius
// 	return diameter, circumference
// }

// func threeDigitNum(number int) (int, int, int){
// 	units := number % 10
// 	tens := number % 100 / 10
// 	hundredths := number / 100
// 	return units, tens, hundredths
// }

// func simpleDigits(N int) {
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

