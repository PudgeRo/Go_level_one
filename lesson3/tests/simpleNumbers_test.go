package tests

import (
	"Go_level_one/lesson3/simpleNumbers"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestSimpleNumbers(t *testing.T) {
	testCase := []struct {
		name          string
		input         int
		output        []int
		expectedError error
	}{
		{
			name:   "case for 30",
			input:  30,
			output: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29},
		},
		{
			name:          "case for 1",
			input:         1,
			expectedError: simpleNumbers.NumMoreThanOne,
		},
		{
			name:          "case for 2",
			input:         2,
			output:        []int{2},
			expectedError: nil,
		},
	}
	for _, cse := range testCase {
		cse := cse
		t.Run(cse.name, func(t *testing.T) {
			res, err := simpleNumbers.SimpleNumbers(cse.input)
			if err != nil {
				if !errors.Is(err, cse.expectedError) {
					t.Fatalf("result is not correct, got '%v', expected '%v'", err, cse.expectedError)
				}
				return
			}
			if !reflect.DeepEqual(res, cse.output) {
				t.Fatalf("result is not correct, got %v, expected %v", res, cse.output)
			}
		})
	}
}

func ExampleSimpleNumbers() {
	slSimpleNums, err := simpleNumbers.SimpleNumbers(30)
	if err != nil {
		panic(err)
	}
	fmt.Println(slSimpleNums)

	//Output: [2, 3, 5, 7, 11, 13, 17, 19, 23, 29]
}
