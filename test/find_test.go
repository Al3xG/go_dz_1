package test

import (
	"go_dz1/find"
	"testing"
)

func TestFind1(t *testing.T) {
	haystack := "Строка для теста 123"
	needle := "123"
	idxS, idxE, isFind := find.FindWSL(haystack, needle)

	expIdxS := 17
	expIdxE := 19
	expFind := true
	if idxS != expIdxS || idxE != expIdxE || isFind != expFind {
		t.Errorf("Result was incorrect, got: %d, want: %d.", idxS, expIdxS)
	}
}

func TestFindEmptyNeedle(t *testing.T) {
	haystack := "Строка"
	needle := ""
	idxS, idxE, isFind := find.Find(haystack, needle)

	expIdxS := -1
	expIdxE := 0
	expFind := false
	if idxS != expIdxS || idxE != expIdxE || isFind != expFind {
		t.Errorf("Result was incorrect, got: %d, want: %d.", idxS, expIdxS)
	}
}

func TestFind2(t *testing.T) {
	haystack := "Строка"
	needle := "Стр"
	idxS, idxE, isFind := find.Find(haystack, needle)

	expIdxS := 0
	expIdxE := 2
	expFind := true
	if idxS != expIdxS || idxE != expIdxE || isFind != expFind {
		t.Errorf("Result was incorrect, got: %d, want: %d.", idxS, expIdxS)
	}
}

func TestFind3(t *testing.T) {
	haystack := "недокоты"
	needle := "кот"
	idxS, idxE, isFind := find.Find(haystack, needle)

	expIdxS := 4
	expIdxE := 6
	expFind := true
	if idxS != expIdxS || idxE != expIdxE || isFind != expFind {
		t.Errorf("Result was incorrect, got: %d, want: %d.", idxS, expIdxS)
	}
}

func TestFindWSL1(t *testing.T) {
	haystack := "Строка для теста 123"
	needle := "123"
	idxS, idxE, isFind := find.FindWSL(haystack, needle)

	expIdxS := 17
	expIdxE := 19
	expFind := true
	if idxS != expIdxS || idxE != expIdxE || isFind != expFind {
		t.Errorf("Result was incorrect, got: %d, want: %d.", idxS, expIdxS)
	}
}

func TestFindWSLEmptyNeedle(t *testing.T) {
	haystack := "Строка"
	needle := ""
	idxS, idxE, isFind := find.FindWSL(haystack, needle)

	expIdxS := -1
	expIdxE := 0
	expFind := false
	if idxS != expIdxS || idxE != expIdxE || isFind != expFind {
		t.Errorf("Result was incorrect, got: %d, want: %d.", idxS, expIdxS)
	}
}

func TestFindWSL2(t *testing.T) {
	haystack := "Строка"
	needle := "Стр"
	idxS, idxE, isFind := find.FindWSL(haystack, needle)

	expIdxS := 0
	expIdxE := 2
	expFind := true
	if idxS != expIdxS || idxE != expIdxE || isFind != expFind {
		t.Errorf("Result was incorrect, got: %d, want: %d.", idxS, expIdxS)
	}
}

func TestFindWSL3(t *testing.T) {
	haystack := "недокоты"
	needle := "кот"
	idxS, idxE, isFind := find.FindWSL(haystack, needle)

	expIdxS := 4
	expIdxE := 6
	expFind := true
	if idxS != expIdxS || idxE != expIdxE || isFind != expFind {
		t.Errorf("Result was incorrect, got: %d, want: %d.", idxS, expIdxS)
	}
}
