package find

import (
	"strings"
	"unicode/utf8"
)

func Find(haystack, needle string) (idxStart, idxEnd int, isFind bool) {
	// default result
	idxStart = -1
	idxEnd = 0
	isFind = false
	// len of input params
	lenH := len([]rune(haystack))
	lenN := len([]rune(needle))
	// convert input params to runes
	runesH := []rune(haystack)
	runesN := []rune(needle)

	// empty needle
	if lenN == 0 {
		return idxStart, idxEnd, isFind
	}
	// needle > haystack
	if lenN > lenH {
		return idxStart, idxEnd, isFind
	}

	//fmt.Println(runesN, runesH, lenH, lenN)
	for i := 0; i < lenH; i++ {
		if runesH[i] == runesN[0] && (lenH-i) > lenN-1 {
			//fmt.Printf("match %s\n", string(runesH[i]))
			flag := true
			for j := 0; j <= lenN-1; j++ {
				//fmt.Printf("cycle j=%d\n", j)
				if runesH[i+j] != runesN[j] {
					flag = false
				}
			}
			if flag == true {
				idxStart = i
				idxEnd = i + lenN - 1
				isFind = flag
				return idxStart, idxEnd, isFind
			}

		}
	}

	return idxStart, idxEnd, isFind
}

func FindWSL(haystack, needle string) (idxStart, idxEnd int, isFind bool) {
	if len([]rune(needle)) == 0 {
		return -1, idxEnd, isFind
	}
	idxStart = strings.Index(haystack, needle)
	//fmt.Println(idxStart)
	if idxStart != -1 {
		idxStart = utf8.RuneCountInString(haystack[0:idxStart])
		idxEnd = idxStart + len([]rune(needle)) - 1
		isFind = true
	}
	return idxStart, idxEnd, isFind
}
