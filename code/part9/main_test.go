package main

import "testing"

func TestSum(t *testing.T) {
	tables := []struct {
		num1     int
		num2     int
		expected int
	}{
		{1, 1, 2},
		{5, 2, 7},
		{5, 5, 10},
	}

	for _, table := range tables {
		total := Sum(table.num1, table.num2)
		if total != table.expected {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want:%d.", table.num1, table.num2, total, table.expected)
		}
	}
}