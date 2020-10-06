package maxincreasekeepingskyline

import "testing"

type Case struct {
	grid     [][]int
	expected int
}

var cases = []Case{
	{[][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}, 35},
	{[][]int{}, 0},
}

func TestMaxIncreaseKeepingSkyline(t *testing.T) {
	for _, c := range cases {
		result := maxIncreaseKeepingSkyline(c.grid)
		if result != c.expected {
			t.Errorf("maxIncreaseKeepingSkyline(%v) == %v, want %v",
				c.grid,
				result,
				c.expected,
			)
		}
	}
}
