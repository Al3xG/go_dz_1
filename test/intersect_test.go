package test

import (
	"go_dz1/intersect"
	"reflect"
	"testing"
)

func TestIntersect1(t *testing.T) {
	sl1 := []int{1, 2, 3, 13, 15, 8, 7, 0}
	sl2 := []int{1, 2, 3, 5, 10, 11, 13, 15}
	expect := []int{1, 2, 3, 13, 15}
	result := intersect.Intersect(sl1, sl2)
	if reflect.DeepEqual(expect, result) == false {
		t.Errorf("Result was incorrect, got: %v, want: %v", result, expect)
	}
}
func TestIntersect2(t *testing.T) {
	sl1 := []int{1, 2, 3}
	sl2 := []int{1, 2, 3}
	expect := []int{1, 2, 3}
	result := intersect.Intersect(sl1, sl2)
	if reflect.DeepEqual(expect, result) == false {
		t.Errorf("Result was incorrect, got: %v, want: %v", result, expect)
	}
}
func TestIntersect3(t *testing.T) {
	sl1 := []int{1, 2, 3}
	sl2 := []int{}
	expect := []int{}
	result := intersect.Intersect(sl1, sl2)
	if reflect.DeepEqual(expect, result) == false && (len(expect) != 0 || len(result) != 0) {
		t.Errorf("Result was incorrect, got: %v, want: %v", result, expect)
	}
}
