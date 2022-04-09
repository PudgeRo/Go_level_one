package main

import (
	"fmt"
	"math"
)

func main() {
	rectangle()
	fmt.Println("")
	circle()
	fmt.Println("")
	threeDigitNum()
}

func rectangle() {
	var length, width float64
	fmt.Print("Enter the length: ")
	fmt.Scan(&length)
	fmt.Print("Enter the width: ")
	fmt.Scan(&width)
	rectangleArea := length * width
	fmt.Printf("The area of a rectangle is %.2f\n", rectangleArea)
}

func circle() {
	var circleArea float64
	pi := math.Pi
	fmt.Print("Enter the area of a circle: ")
	fmt.Scan(&circleArea)
	radius := math.Sqrt(circleArea / pi)
	circumference := 2 * pi * radius
	diameter := 2 * radius
	fmt.Printf("The diameter of a circle is %.2f and the circumference is %.2f\n", diameter, circumference)
}

func threeDigitNum() {
	var number int
	fmt.Print("Enter a three-digit number: ")
	fmt.Scan(&number)
	units := number % 10
	tens := number % 100 / 10
	hundredths := number / 100
	fmt.Printf("Number %d\nunits: %d\ntens: %d\nhundredths: %d\n", number, units, tens, hundredths)
}


