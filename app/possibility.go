package main

import (
	"sort"
	"strconv"
)

func GetPossibilities(fieldValue byte, fieldCount int) [][]int {
	var possibilities [][]int
	if fieldCount < 2 {
		return possibilities
	}

	if fieldCount == 2 {
		return getTwoFieldPossibility(fieldValue)
	}
	var checkedPossibilities = map[string]bool{}
	for j := 1; j < 10; j++ {
		p2 := GetPossibilities(byte(int(fieldValue)-j), fieldCount-1)
		for _, v := range p2 {
			if !contains(v, j) {
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
	}

	for _, v := range possibilities {
		sort.Ints(v)
	}

	return possibilities
}
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func getTwoFieldPossibility(fieldValue byte) (possibilities [][]int) {
	possibilities = [][]int{}
	var checkedPossibilities = map[string]bool{}
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i != j {
				if i+j == int(fieldValue) {
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
