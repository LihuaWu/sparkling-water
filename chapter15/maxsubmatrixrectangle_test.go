package chapter15

import "testing"

var maxSubmatrixRectangleTests = []struct {
	mat      [][]bool
	expected int
}{
	{
		[][]bool{
			{true, true, false},
			{true, true, true},
		},
		4,
	},
	{
		[][]bool{
			{true, true, false, false},
			{true, true, true, true},
			{true, true, true, true},
		},
		8,
	},
}

func TestMaxSubmatrixRectangle(t *testing.T) {
	for i, tt := range maxSubmatrixRectangleTests {
		actual := MaxSubmatrixRectangle(tt.mat)
		if actual != tt.expected {
			t.Errorf("MaxSubmatrixRectangle [%d]: expected %d, actual %d", i+1, tt.expected, actual)
		}
	}
}