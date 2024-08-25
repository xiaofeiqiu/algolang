package main

import (
	"reflect"
	"testing"
)

// The delete function should be defined here or imported

func TestDelete(t *testing.T) {
	tests := []struct {
		name  string
		cover []Block
		index int
		want  []Block
	}{
		{
			name:  "Delete from middle of first block",
			cover: []Block{{7, 11}, {3, 6}},
			index: 1,
			want:  []Block{{7, 8}, {9, 11}, {3, 6}},
		},
		{
			name:  "Delete from end of first block",
			cover: []Block{{7, 11}, {3, 6}},
			index: 3,
			want:  []Block{{7, 10}, {3, 6}},
		},
		{
			name:  "Delete from start of second block",
			cover: []Block{{7, 11}, {3, 6}},
			index: 4,
			want:  []Block{{7, 11}, {4, 6}},
		},
		{
			name:  "Delete from middle of second block",
			cover: []Block{{7, 11}, {3, 6}},
			index: 5,
			want:  []Block{{7, 11}, {3, 4}, {5, 6}},
		},
		{
			name:  "Delete after all blocks",
			cover: []Block{{7, 11}, {3, 6}},
			index: 7,
			want:  []Block{{7, 11}, {3, 6}},
		},
		{
			name:  "Delete from single-character block",
			cover: []Block{{7, 8}, {3, 6}},
			index: 0,
			want:  []Block{{3, 6}},
		},
		{
			name:  "Delete from empty cover",
			cover: []Block{},
			index: 0,
			want:  []Block{},
		},
		{
			name:  "Delete from cover with single block",
			cover: []Block{{3, 6}},
			index: 1,
			want:  []Block{{3, 4}, {5, 6}},
		},
		{
			name:  "Delete last character",
			cover: []Block{{7, 11}, {3, 6}},
			index: 6,
			want:  []Block{{7, 11}, {3, 5}},
		},
		{
			name:  "Delete single-character block in the middle",
			cover: []Block{{1, 3}, {5, 6}, {7, 10}},
			index: 2,
			want:  []Block{{1, 3}, {7, 10}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := delete1(tt.cover, tt.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

// You can add more specific tests here if needed
