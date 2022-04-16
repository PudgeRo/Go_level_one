package main

import (
	"fmt"
	"strconv"
)

func main() {
	var slice []int
	var number string
	fmt.Println("Enter 'stop' to stop program")
	for {
		fmt.Print("Enter number: ")
		fmt.Scan(&number)
		if number == "stop" {
			break
		}
		number, _ := strconv.Atoi(number)
		slice = append(slice, number)
	}
	slice = sortByInsert(slice)
	fmt.Printf("Sorted array: %v\n", slice)

	// t1 := []int{7, 4, 3, 8, 5, 2, 0, 1}
	// fmt.Println(sortByInsert(t1))
	// t2 := []int{8, 8, 8, 8, 8, 8, 8}
	// fmt.Println(sortByInsert(t2))
	// t3 := []int{}
	// fmt.Println(sortByInsert(t3))
	// t4 := []int{-4, -1, -6, -8, -2}
	// fmt.Println(sortByInsert(t4))
}

func sortByInsert(slice []int) []int {
	var x int
	for i := 1; i < len(slice); i++ {
		x = slice[i]
		for (i > 0 && slice[i-1] > x) {
			slice[i] = slice[i - 1]
			i--
		}
		slice[i] = x 
	}
	return slice
}