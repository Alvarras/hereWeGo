package challenge

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type testCase struct {
		nums     []int
		target   int
		expected []int
	}

	tests := []testCase{
		{[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			19,
			[]int{8, 9},
		},
	}

	passCount := 0
	failCount := 0

	for i, tc := range tests {
		result := twoSum(tc.nums, tc.target)
		if reflect.DeepEqual(result, tc.expected) {
			passCount++
			t.Logf("Test case %d PASSED:\nnums: %v\ntarget: %d\nactual: %v\n", i+1, tc.nums, tc.target, result)
		} else {
			failCount++
			t.Errorf("Test case %d FAILED:\nnums: %v\ntarget: %d\nexpected: %v\nactual: %v\n\n", i+1, tc.nums, tc.target, tc.expected, result)
		}
	}

	t.Logf("\nTotal Passed tests: %d", passCount)
	t.Logf("\nTotal Failed tests: %d", failCount)

}
