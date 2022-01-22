package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Possibilities struct {
	value        int
	fieldCount   int
	combinations [][]int
}

func GetPossibilities(fieldValue int, fieldCount int, knownCombinations map[string][][]int) Possibilities {
	possibilities := Possibilities{
		value:      fieldValue,
		fieldCount: fieldCount,
	}

	if fieldCount < 1 || fieldValue < 1 {
		return possibilities
	}

	if fieldCount == 1 {
		possibilities.combinations = [][]int{{fieldValue}}
	}

	if fieldCount == 2 {
		return getTwoFieldPossibilities(fieldValue)
	}

	var checkedPossibilities = map[string]bool{}
	for j := 1; j < 10; j++ {
		combinationKey := fmt.Sprintf("%d;%d", int(fieldValue)-j, fieldCount-1)
		val, ok := knownCombinations[combinationKey]
		if !ok {
			val = GetPossibilities(fieldValue-j, fieldCount-1, knownCombinations).combinations
			knownCombinations[combinationKey] = val
		}
		for _, v := range val {
			if sum(v)+j != fieldValue {
				continue
			}
			if contains(v, j) {
				continue
			}

			v = append(v, j)
			key := generateCombinationKey(v)
			if !checkedPossibilities[key] {
				checkedPossibilities[key] = true
				possibilities.combinations = append(possibilities.combinations, v)
			}
		}
	}

	for _, v := range possibilities.combinations {
		sort.Ints(v)
	}

	return possibilities
}

func generateCombinationKey(possibilities []int) string {
	sort.Ints(possibilities)
	var key string
	for _, w := range possibilities {
		key = key + strconv.Itoa(w)
	}

	return key
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getTwoFieldPossibilities(fieldValue int) Possibilities {
	possibilities := Possibilities{
		value:      fieldValue,
		fieldCount: 2,
	}
	var checkedPossibilities = map[string]bool{}
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i == j {
				continue
			}

			if i+j != fieldValue {
				continue
			}

			possibility := []int{i, j}
			key := generateCombinationKey(possibility)

			if !checkedPossibilities[key] {
				checkedPossibilities[key] = true
				possibilities.combinations = append(possibilities.combinations, possibility)
			}
		}
	}

	return possibilities
}
