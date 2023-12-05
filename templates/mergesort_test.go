package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Normal Case", []int{38, 27, 43, 3, 9, 82, 10}, []int{3, 9, 10, 27, 38, 43, 82}},
		{"Empty Array", []int{}, []int{}},
		{"Single Element", []int{1}, []int{1}},
		{"Two Elements", []int{2, 1}, []int{1, 2}},
		{"Negative Numbers", []int{-3, -1, -2}, []int{-3, -2, -1}},
		{"Array with Duplicates", []int{5, 1, 5, 3}, []int{1, 3, 5, 5}},
		{"Large Array", []int{100, 99, 98, 1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5, 98, 99, 100}},
		{"Sorted Array", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"Reverse Sorted Array", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := mergeSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("For %s, expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
