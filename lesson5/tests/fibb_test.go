package tests

import (
	"Go_level_one/lesson5/fibb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci(t *testing.T) {
	testCase := []struct {
		input  int
		output int
	}{
		{
			input:  6,
			output: 8,
		},
		{
			input:  2,
			output: 1,
		},
		{
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
		res, _ := fibb.Fibonacci(cse.input)
		assert.Equal(t, cse.output, res)
	}

	for _, cse := range testCaseWithErrors {
		_, err := fibb.Fibonacci(cse.input)
		assert.EqualError(t, err, "number must be positive")
	}
}
