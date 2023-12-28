package main

import "testing"

func TestNumberofLakes(t *testing.T) {
	tcs := []struct {
		name     string
		grid     [][]rune
		expected int
	}{
		{
			name: "One Big Lake",
			grid: [][]rune{
				{'x', 'x', 'x', 'x'},
				{'x', '.', '.', 'x'},
				{'x', '.', '.', 'x'},
				{'x', 'x', 'x', 'x'},
			},
			expected: 1,
		},
		{
			name: "One Big Lake",
			grid: [][]rune{
				{'x', 'x', 'x', 'x'},
				{'x', '.', '.', 'x'},
				{'x', '.', '.', 'x'},
				{'x', 'x', 'x', 'x'},
				{'x', 'x', 'x', 'x'},
				{'x', 'x', 'x', 'x'},
				{'x', '.', 'x', 'x'},
				{'x', 'x', 'x', 'x'},
			},
			expected: 2,
		},
		{
			name: "One Big Lake",
			grid: [][]rune{
				{'x', 'x', 'x', 'x'},
				{'x', '.', '.', 'x'},
				{'x', '.', '.', 'x'},
				{'x', 'x', 'x', 'x'},
				{'x', 'x', 'x', 'x'},
				{'.', 'x', '.', 'x'},
				{'x', '.', 'x', 'x'},
				{'x', 'x', 'x', 'x'},
			},
			expected: 3,
		},
	}

	for _, tc := range tcs {
		act := countLakes(tc.grid)
		if act != tc.expected {
			t.Errorf("want %d, got %d", tc.expected, act)
		}
	}

}
