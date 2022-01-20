package main

import "testing"

type TestSet struct {
	value        int
	fieldCount   int
	expectations [][]int
}

func TestGetPossibilities(t *testing.T) {
	var testSets []TestSet

	testSets = append(testSets, TestSet{
		value:        3,
		fieldCount:   2,
		expectations: [][]int{{1, 2}},
	})

	testSets = append(testSets, TestSet{
		value:        4,
		fieldCount:   2,
		expectations: [][]int{{1, 3}},
	})

	testSets = append(testSets, TestSet{
		value:        5,
		fieldCount:   2,
		expectations: [][]int{{1, 4}, {2, 3}},
	})

	testSets = append(testSets, TestSet{
		value:        17,
		fieldCount:   2,
		expectations: [][]int{{8, 9}},
	})

	testSets = append(testSets, TestSet{
		value:        2,
		fieldCount:   2,
		expectations: [][]int{},
	})

	testSets = append(testSets, TestSet{
		value:        18,
		fieldCount:   2,
		expectations: [][]int{},
	})
	testSets = append(testSets, TestSet{
		value:        6,
		fieldCount:   3,
		expectations: [][]int{{1, 2, 3}},
	})
	testSets = append(testSets, TestSet{
		value:        8,
		fieldCount:   3,
		expectations: [][]int{{1, 2, 5}, {1, 3, 4}},
	})
	testSets = append(testSets, TestSet{
		value:        10,
		fieldCount:   4,
		expectations: [][]int{{1, 2, 3, 4}},
	})
	testSets = append(testSets, TestSet{
		value:        45,
		fieldCount:   9,
		expectations: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	})

	for _, testSet := range testSets {
		possibilities := GetPossibilities(byte(testSet.value), testSet.fieldCount)
		if len(possibilities) != len(testSet.expectations) {
			t.Errorf("Expected possibilities to be %d instead got %d.", len(testSet.expectations), len(possibilities))
			t.Error(testSet.expectations, possibilities)
			return
		}

		for i, p := range possibilities {
			if !equal(p, testSet.expectations[i]) {
				t.Error("Possibilities don't match expectations", p, testSet.expectations[i])
			}
			if sum(p) != testSet.value {
				t.Error("Sum is not correct", sum(p), p, testSet.value)
			}
		}

	}
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
