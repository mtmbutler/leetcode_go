package subrectanglequeries

import "testing"

func TestSubrectangleQueries(t *testing.T) {
	arr := [][]int{
		{1, 2, 1},
		{4, 3, 4},
		{3, 2, 1},
		{1, 1, 1},
	}
	sq := Constructor(arr)

	// Test GetValue
	v1 := sq.GetValue(0, 2)
	expected := 1
	if v1 != expected {
		t.Errorf("subrectangleQueries.GetValue(0, 2) == %v, want %v", v1, expected)
	}

	// Test more GetValues after updating
	sq.UpdateSubrectangle(0, 0, 3, 2, 5)
	v1 = sq.GetValue(0, 2)
	expected = 5
	if v1 != expected {
		t.Errorf("subrectangleQueries.GetValue(0, 2) == %v, want %v", v1, expected)
	}
	v1 = sq.GetValue(3, 1)
	expected = 5
	if v1 != expected {
		t.Errorf("subrectangleQueries.GetValue(3, 1) == %v, want %v", v1, expected)
	}

	// Test more GetValues after updating
	sq.UpdateSubrectangle(3, 0, 3, 2, 10)
	v1 = sq.GetValue(0, 2)
	expected = 5
	if v1 != expected {
		t.Errorf("subrectangleQueries.GetValue(0, 2) == %v, want %v", v1, expected)
	}
	v1 = sq.GetValue(3, 1)
	expected = 10
	if v1 != expected {
		t.Errorf("subrectangleQueries.GetValue(3, 1) == %v, want %v", v1, expected)
	}
}
