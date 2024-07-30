package rmdpl

import "slices"

func RemoveDuplicates(input []int) []int {
	var result []int
	for _, v := range input {
		flag := false
		for _, v2 := range result {
			if v == v2 {
				flag = true
			}
		}
		if flag == false {
			result = append(result, v)
		}
	}
	return result
}

func RemoveDuplicates2(input []int) []int {
	slices.Sort(input)
	return slices.Compact(input)
}
