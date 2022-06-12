package tests

import (
	"Go_level_one/lesson4/sort"
	"reflect"
	"testing"
)

type testCaseForSortByInsert struct {
	name   string
	input  []int
	output []int
}

func TestSortByInsert(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "case for [1, 3, 6, 2]",
			input:  []int{1, 3, 6, 2},
			output: []int{1, 2, 3, 6},
		},
		{
			name:   "case for [-2, -8, 0, -54, 4, 1]",
			input:  []int{-2, -8, 0, -54, 4, 1},
			output: []int{-54, -8, -2, 0, 1, 4},
		},
		{
			name:   "case for []",
			input:  []int{},
			output: []int{},
		},
	}
	for _, cse := range testCases {
		cse := cse
		t.Run(cse.name, func(t *testing.T) {
			res := sort.SortByInsert(cse.input)
			if !reflect.DeepEqual(res, cse.output) {
				t.Fatalf("result is not correct, got %v, expected %v", res, cse.output)
			}
		})
	}
}

func BenchmarkSortByInsert(b *testing.B) {
	testSlice := []int{3, 5, 1, 6, 87, 321, -132, 34, 0, -4}
	for i := 0; i < b.N; i++ {
		_ = sort.SortByInsert(testSlice)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	testSlice := []int{3, 5, 1, 6, 87, 321, -132, 34, 0, -4}
	for i := 0; i < b.N; i++ {
		_ = sort.BubleSort(testSlice)
	}
}
