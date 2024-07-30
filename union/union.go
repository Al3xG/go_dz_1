package union

func Union(slice1 []int, slice2 []int) []int {
	result := make([]int, len(slice1)+len(slice2))
	copy(result, slice1)
	copy(result[len(slice1):], slice2)
	return result
}
