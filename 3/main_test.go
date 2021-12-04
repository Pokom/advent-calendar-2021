package main

import "testing"

func TestBinaryToInt(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{
			name:   "Base Case 0",
			input:  []int{0},
			output: 0,
		},
		{
			name:   "Base Case 1",
			input:  []int{1},
			output: 1,
		},
		{
			name:   "Base Case 2",
			input:  []int{0, 0},
			output: 0,
		},
		{
			name:   "Base Case 3",
			input:  []int{0, 1},
			output: 1,
		},
		{
			name:   "Base Case 4",
			input:  []int{1, 0},
			output: 2,
		},
		{
			name:   "Base Case 4",
			input:  []int{1, 1},
			output: 3,
		},
		{
			name:   "Base Case 5",
			input:  []int{1, 0, 0},
			output: 4,
		},
		{
			name:   "Base Case 5",
			input:  []int{1, 0, 1},
			output: 5,
		},
		{
			name:   "Base Case 5",
			input:  []int{1, 1, 0},
			output: 6,
		},
		{
			name:   "Base Case 5",
			input:  []int{1, 1, 1},
			output: 7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := binaryToInt(tc.input)
			if result != tc.output {
				t.Errorf("Expected %d, got %d instead\n", tc.output, result)
			}
		})
	}
}
