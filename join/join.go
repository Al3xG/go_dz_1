package join

import "strings"

func Join(str []rune) string {
	result := ""
	for el := range str {
		result += string(str[el])
	}
	return result
}

func Join2(str []rune) string {
	result := make([]string, 0, len(str))
	for el := range str {
		result = append(result, string(str[el]))
	}
	return strings.Join(result, "")
}
