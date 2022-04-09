package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// //The area of the rectangle
	var length, width float64
	fmt.Print("Enter the length: ")
	fmt.Scan(&length)
	fmt.Print("Enter the width: ")
	fmt.Scan(&width)
	fmt.Printf("The area of the rectangle is %.2f\n\n", rectangle(length, width))

	// //Diameter and circumference 
	var circleArea float64
	fmt.Print("Enter the area of a circle: ")
	fmt.Scan(&circleArea)
	diameter, circumference := circle(circleArea)
	fmt.Printf("Diameter is %.2f, circumference is %.2f\n\n", diameter, circumference)

	// //Decomposition of number
	var number int
	fmt.Print("Enter a three-digit number: ")
	fmt.Scan(&number)
	units, tens, hundredths := threeDigitNum(number)
	fmt.Printf("Units: %d\nTens: %d\nHundredths: %d\n\n", units, tens, hundredths)

	// //Find simple number from 0 to N
	var N int
	fmt.Print("Enter N: ")
	fmt.Scan(&N)
	simpleDigits(N)

	//Calculator 
	var number1, number2 string
	var operation string
	fmt.Println("Enter 'stop' instead of first number for stopping program")
	for {
		fmt.Print("Enter first number: ")
		fmt.Scan(&number1)
		if number1 == "stop" {
			break
		}
		fmt.Print("Enter second number: ")
		fmt.Scan(&number2)
		fmt.Print("Enter operation: ")
		fmt.Scan(&operation)
		firstNum, _ := strconv.ParseFloat(number1, 64)
		secondNum, _ := strconv.ParseFloat(number2, 64)
		calculator(firstNum, secondNum, operation)
	}
}

func rectangle(length, width float64) float64{
	rectangleArea := length * width
	return rectangleArea
}

func circle(circleArea float64) (float64, float64){
	pi := math.Pi
	radius := math.Sqrt(circleArea / pi)
	circumference := 2 * pi * radius
	diameter := 2 * radius
	return diameter, circumference
}

func threeDigitNum(number int) (int, int, int){
	units := number % 10
	tens := number % 100 / 10
	hundredths := number / 100
	return units, tens, hundredths
}

func simpleDigits(N int) {
	var isSimple bool
	for i := 1; i <= N; i++ {
		isSimple = true
		for j := 2; j < i; j++ {
			if i % j == 0 {
				isSimple = false
				break
			}
		}
		if isSimple {
			fmt.Printf("%d- простое число\n", i)
		}
	}
}

func calculator(number1, number2 float64, operation string) {
	if operation == "+" {
		fmt.Printf("Result %.2f\n", number1 + number2)
	} else if operation == "-" {
		fmt.Printf("Result %.2f\n", number1 - number2)
	} else if operation == "*" {
		fmt.Printf("Result %.2f\n", number1 * number2)
	} else if operation == "/" {
		if number2 == 0 {
			fmt.Printf("Division by zero\n")
		} else {
			fmt.Printf("Result %.2f\n", number1 / number2)
		}
	}
}