package intersect

func Intersect(slice1 []int, slice2 []int) []int {
	var result []int
	for _, i := range slice1 {
		for _, i2 := range slice2 {
			if i == i2 {
				result = append(result, i)
			}
		}
	}
	return result
}
