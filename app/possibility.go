package main

import (
	"fmt"
	"sort"
	"strconv"
)

func GetPossibilities(fieldValue int, fieldCount int, knownCombinations map[string][][]int) [][]int {
	var possibilities [][]int
	if fieldCount < 2 {
		return possibilities
	}

	if fieldCount == 2 {
		return getTwoFieldPossibility(fieldValue)
	}

	var checkedPossibilities = map[string]bool{}
	for j := 1; j < 10; j++ {
		combinationKey := fmt.Sprintf("%d;%d", int(fieldValue)-j, fieldCount-1)
		val, ok := knownCombinations[combinationKey]
		if !ok {
			val = GetPossibilities(fieldValue-j, fieldCount-1, knownCombinations)
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
			sort.Ints(v)
			var key string
			for _, w := range v {
				key = key + strconv.Itoa(w)
			}
			if !checkedPossibilities[key] {
				checkedPossibilities[key] = true
				possibilities = append(possibilities, v)
			}
		}
	}

	for _, v := range possibilities {
		sort.Ints(v)
	}

	return possibilities
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
func getTwoFieldPossibility(fieldValue int) [][]int {
	var possibilities [][]int
	var checkedPossibilities = map[string]bool{}
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i != j {
				if i+j == fieldValue {
					possibility := []int{i, j}
					sort.Ints(possibility)
					var key string
					for _, v := range possibility {
						key = key + strconv.Itoa(v)
					}

					if !checkedPossibilities[key] {
						checkedPossibilities[key] = true
						possibilities = append(possibilities, possibility)
					}
				}
			}
		}
	}

	return possibilities
}
