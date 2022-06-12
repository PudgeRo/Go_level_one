package tests

import (
	"Go_level_one/lesson5/fibb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci(t *testing.T) {
	testCase := []struct {
		name   string
		input  int
		output int
	}{
		{
			name:   "case with six",
			input:  6,
			output: 8,
		},
		{
			name:   "case with two",
			input:  2,
			output: 1,
		},
		{
			name:   "case with zero",
			input:  0,
			output: 0,
		},
	}
	testCaseWithErrors := []struct {
		input         int
		expectedError error
	}{
		{
			input:         -1,
			expectedError: fibb.ErrNonPositive,
		},
	}

	for _, cse := range testCase {
		t.Run(cse.name, func(t *testing.T) {
			res, _ := fibb.Fibonacci(cse.input)
			assert.Equal(t, cse.output, res)
		})
	}

	for _, cse := range testCaseWithErrors {
		_, err := fibb.Fibonacci(cse.input)
		assert.EqualError(t, err, "number must be positive")
	}
}
