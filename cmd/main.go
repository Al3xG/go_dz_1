package main

import (
	"fmt"
	"go_dz1/find"
	"go_dz1/intersect"
	"go_dz1/join"
	"go_dz1/reverse"
	"go_dz1/rmdpl"
	"go_dz1/union"
)

func main() {
	// reverse block
	revTest := "Это пример Reverse!"
	fmt.Println("#1", reverse.Reverse(revTest))
	fmt.Println("#2", reverse.Reverse2(revTest))
	fmt.Println("#3", reverse.Reverse3(revTest))
	fmt.Println()
	// find block
	haystack := "Тестовая строка, а тут кот, продолжение строки"
	needle := "кот"
	idxStart, idxEnd, isFind := find.Find(haystack, needle)
	fmt.Println("#1 w/o ext lib:")
	fmt.Printf("idxStart: %d, idxEnd: %d, isFind:%t\n", idxStart, idxEnd, isFind)
	idxStart, idxEnd, isFind = find.FindWSL(haystack, needle)
	fmt.Println("#2 with ext lib:")
	fmt.Printf("idxStart: %d, idxEnd: %d, isFind:%t\n", idxStart, idxEnd, isFind)
	// join runes
	runes := []rune(revTest)
	fmt.Println()
	fmt.Println("#1 w/o ext lib:")
	fmt.Println(join.Join(runes[:10]))
	fmt.Println("#2 with ext lib:")
	fmt.Println(join.Join2(runes[:10]))
	fmt.Println()
	// remove duplicates from slice of int
	inpData := []int{1, 2, 2, 2, 12, 3, 2, 2, 1, 33, 4, 5, 6}
	fmt.Println(inpData)
	fmt.Println("Remove duplicates ver1: ", rmdpl.RemoveDuplicates(inpData))
	fmt.Println("Remove duplicates ver2: ", rmdpl.RemoveDuplicates2(inpData))
	fmt.Println()
	// intersection of slices
	sl1 := []int{1, 2, 3, 13, 15, 8, 7, 0}
	sl2 := []int{1, 2, 3, 5, 10, 11, 13, 15}
	fmt.Println("SL1:", sl1)
	fmt.Println("SL2:", sl2)
	fmt.Println("Intersection of sl1 and sl2:", intersect.Intersect(sl1, sl2))
	fmt.Println()
	// union of two slices
	fmt.Println("SL1:", sl1)
	fmt.Println("SL2:", sl2)
	fmt.Println("Union of sl1 and sl2:", union.Union(sl1, sl2))

}
