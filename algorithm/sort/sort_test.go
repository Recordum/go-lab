package sort

import (
	"reflect"
	"testing"
)

// 테스트 케이스 구조체
type testCase struct {
	input    []int
	expected []int
}

// 테스트 케이스 목록
var testCases = []testCase{
	{input: []int{5, 2, 9, 1, 5, 6}, expected: []int{1, 2, 5, 5, 6, 9}},
	{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
	{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
	{input: []int{1, 1, 1, 1, 1}, expected: []int{1, 1, 1, 1, 1}},
	{input: []int{}, expected: []int{}},
	{input: []int{42}, expected: []int{42}},
}

func TestBubbleSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run("BubbleSort", func(t *testing.T) {
			result := BubbleSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("BubbleSort failed: got %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run("SelectionSort", func(t *testing.T) {
			result := SelectionSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("SelectionSort failed: got %v, want %v", result, tc.expected)
			}
		})
	}
}

// func TestInsertionSort(t *testing.T) {
// 	for _, tc := range testCases {
// 		t.Run("InsertionSort", func(t *testing.T) {
// 			result := InsertionSort(tc.input)
// 			if !reflect.DeepEqual(result, tc.expected) {
// 				t.Errorf("InsertionSort failed: got %v, want %v", result, tc.expected)
// 			}
// 		})
// 	}
// }
