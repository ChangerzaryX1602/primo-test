package usecase_test

import (
	"reflect"
	"testing"

	"test/internal/repository"
	"test/internal/usecase"
)

type TestCase struct {
	Name     string
	Col1     []int64
	Col2     []int64
	Col3     []int64
	Expected []int64
}

func TestMerge(t *testing.T) {
	repo := repository.NewTestRepository(nil)
	uc := usecase.NewTestUsecase(repo)

	testCases := []TestCase{
		{
			Name:     "TestCase 1: Normal merge",
			Col1:     []int64{1, 3, 5, 7},
			Col2:     []int64{2, 4, 6, 8},
			Col3:     []int64{20, 15, 10, 0},
			Expected: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 10, 15, 20},
		},
		{
			Name:     "TestCase 2: Empty col1",
			Col1:     []int64{},
			Col2:     []int64{1, 2, 3},
			Col3:     []int64{6, 5, 4},
			Expected: []int64{1, 2, 3, 4, 5, 6},
		},
		{
			Name:     "TestCase 3: Empty col2",
			Col1:     []int64{1, 2, 3},
			Col2:     []int64{},
			Col3:     []int64{9, 7, 5},
			Expected: []int64{1, 2, 3, 5, 7, 9},
		},
		{
			Name:     "TestCase 4: Empty col3",
			Col1:     []int64{1, 2, 3},
			Col2:     []int64{4, 5, 6},
			Col3:     []int64{},
			Expected: []int64{1, 2, 3, 4, 5, 6},
		},
		{
			Name:     "TestCase 5: All empty slices",
			Col1:     []int64{},
			Col2:     []int64{},
			Col3:     []int64{},
			Expected: []int64{},
		},
		{
			Name:     "TestCase 6: Single element arrays",
			Col1:     []int64{1},
			Col2:     []int64{2},
			Col3:     []int64{3},
			Expected: []int64{1, 2, 3},
		},
		{
			Name:     "TestCase 7: With duplicates",
			Col1:     []int64{1, 3, 3, 5},
			Col2:     []int64{2, 3, 6},
			Col3:     []int64{10, 8, 8, 7},
			Expected: []int64{1, 2, 3, 3, 3, 5, 6, 7, 8, 8, 10},
		},
		{
			Name:     "TestCase 8: Negative numbers",
			Col1:     []int64{-10, -5, 0},
			Col2:     []int64{-8, -3, 1},
			Col3:     []int64{5, 2, -2},
			Expected: []int64{-10, -8, -5, -3, -2, 0, 1, 2, 5},
		},
		{
			Name:     "TestCase 9: Mixed numbers",
			Col1:     []int64{0, 5, 10},
			Col2:     []int64{3, 8, 12},
			Col3:     []int64{20, 15, 11, -1},
			Expected: []int64{-1, 0, 3, 5, 8, 10, 11, 12, 15, 20},
		},
		{
			Name:     "TestCase 10: All same numbers",
			Col1:     []int64{5, 5, 5},
			Col2:     []int64{5, 5},
			Col3:     []int64{5, 5, 5, 5},
			Expected: []int64{5, 5, 5, 5, 5, 5, 5, 5, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := uc.Merge(tc.Col1, tc.Col2, tc.Col3)
			if !reflect.DeepEqual(result, tc.Expected) {
				t.Errorf("%s: Expected %v, got %v", tc.Name, tc.Expected, result)
			}
			if err != nil {
				t.Errorf("%s: Unexpected error: %v", tc.Name, err)
			}
		})
	}
}

type MergeTwoTestCase struct {
	Name     string
	A        []int64
	B        []int64
	Expected []int64
}

func TestMergeTwo(t *testing.T) {
	testCases := []MergeTwoTestCase{
		{
			Name:     "TestCase 1: Normal mergeTwo",
			A:        []int64{1, 4, 7},
			B:        []int64{2, 3, 8, 9},
			Expected: []int64{1, 2, 3, 4, 7, 8, 9},
		},
		{
			Name:     "TestCase 2: First slice empty",
			A:        []int64{},
			B:        []int64{1, 2, 3},
			Expected: []int64{1, 2, 3},
		},
		{
			Name:     "TestCase 3: Second slice empty",
			A:        []int64{1, 2, 3},
			B:        []int64{},
			Expected: []int64{1, 2, 3},
		},
		{
			Name:     "TestCase 4: Both slices empty",
			A:        []int64{},
			B:        []int64{},
			Expected: []int64{},
		},
		{
			Name:     "TestCase 5: Non-overlapping ranges",
			A:        []int64{1, 2, 3},
			B:        []int64{4, 5, 6},
			Expected: []int64{1, 2, 3, 4, 5, 6},
		},
		{
			Name:     "TestCase 6: Reversed order inputs",
			A:        []int64{2, 4, 6},
			B:        []int64{1, 3, 5},
			Expected: []int64{1, 2, 3, 4, 5, 6},
		},
		{
			Name:     "TestCase 7: Alternate merge",
			A:        []int64{1, 3, 5},
			B:        []int64{2, 4, 6},
			Expected: []int64{1, 2, 3, 4, 5, 6},
		},
		{
			Name:     "TestCase 8: With duplicates",
			A:        []int64{1, 2, 2, 3},
			B:        []int64{2, 2, 4},
			Expected: []int64{1, 2, 2, 2, 2, 3, 4},
		},
		{
			Name:     "TestCase 9: Negative numbers",
			A:        []int64{-5, 0, 5},
			B:        []int64{-10, -1, 10},
			Expected: []int64{-10, -5, -1, 0, 5, 10},
		},
		{
			Name:     "TestCase 10: Single element each",
			A:        []int64{1},
			B:        []int64{2},
			Expected: []int64{1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := usecase.MergeTwo(tc.A, tc.B)
			if !reflect.DeepEqual(result, tc.Expected) {
				t.Errorf("%s: Expected %v, got %v", tc.Name, tc.Expected, result)
			}
		})
	}
}

type ReverseTestCase struct {
	Name     string
	Input    []int64
	Expected []int64
}

func TestReverse(t *testing.T) {
	testCases := []ReverseTestCase{
		{
			Name:     "TestCase 1: Normal reverse",
			Input:    []int64{1, 2, 3, 4, 5},
			Expected: []int64{5, 4, 3, 2, 1},
		},
		{
			Name:     "TestCase 2: Empty slice",
			Input:    []int64{},
			Expected: []int64{},
		},
		{
			Name:     "TestCase 3: Single element",
			Input:    []int64{1},
			Expected: []int64{1},
		},
		{
			Name:     "TestCase 4: Even number of elements",
			Input:    []int64{2, 4, 6, 8},
			Expected: []int64{8, 6, 4, 2},
		},
		{
			Name:     "TestCase 5: Descending input",
			Input:    []int64{10, 9, 8, 7, 6, 5},
			Expected: []int64{5, 6, 7, 8, 9, 10},
		},
		{
			Name:     "TestCase 6: Negative numbers",
			Input:    []int64{-1, -2, -3},
			Expected: []int64{-3, -2, -1},
		},
		{
			Name:     "TestCase 7: All zeros",
			Input:    []int64{0, 0, 0},
			Expected: []int64{0, 0, 0},
		},
		{
			Name:     "TestCase 8: With duplicates",
			Input:    []int64{1, 3, 3, 7},
			Expected: []int64{7, 3, 3, 1},
		},
		{
			Name:     "TestCase 9: Mixed positive and negative",
			Input:    []int64{100, 50, 0, -50, -100},
			Expected: []int64{-100, -50, 0, 50, 100},
		},
		{
			Name:     "TestCase 10: Multiple duplicates",
			Input:    []int64{2, 2, 2, 3, 3},
			Expected: []int64{3, 3, 2, 2, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := usecase.Reverse(tc.Input)
			if !reflect.DeepEqual(result, tc.Expected) {
				t.Errorf("%s: Expected %v, got %v", tc.Name, tc.Expected, result)
			}
		})
	}
}
