package diagsort

import (
	"reflect"
	"testing"
)

type Case struct {
	mat      [][]int
	expected [][]int
}

var cases = []Case{
	{
		[][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}},
		[][]int{{1, 1, 1, 1}, {1, 2, 2, 2}, {1, 2, 3, 3}},
	},
}

func TestDiagonalSort(t *testing.T) {
	for _, c := range cases {
		result := diagonalSort(c.mat)
		bothEmpty := len(c.mat) == 0 && len(result) == 0
		if !bothEmpty && !reflect.DeepEqual(result, c.expected) {
			t.Errorf(
				"diagonalSort(%v) == %v, want %v",
				c.mat, result, c.expected,
			)
		}
	}
}
