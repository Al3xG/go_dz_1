package reverse

import "strings"

func Reverse(in string) string {

	workStr := []rune(in)
	result := make([]rune, 0, len(workStr))
	startLen := len(workStr)
	for i := startLen - 1; i >= 0; i-- {
		//fmt.Println("LEN:", len(result), "CAP:", cap(result))
		result = append(result, workStr[i])
	}

	return string(result)
}

func Reverse2(in string) string {
	var result strings.Builder
	workStr := []rune(in)
	startLen := len(workStr)
	for i := startLen - 1; i >= 0; i-- {
		//fmt.Println("LEN:", len(result), "CAP:", cap(result))
		result.WriteRune(workStr[i])
	}

	return result.String()
}

func Reverse3(in string) string {
	result := []rune(in)
	for i, j := 0, len(result)-1; i < len(result)/2; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
