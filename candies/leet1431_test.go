package candies

import (
	"reflect"
	"testing"
)

func TestKidsWithCandiesShouldDisplay4True(t *testing.T) {
	result := kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	if !reflect.DeepEqual(result, []bool{true, true, true, false, true}) {
		t.Errorf("kids with candies should display [true, true, true, false, true] but got %v", result)
	}
}

func TestKidsWithCandiesShouldDisplay1TrueAnd4False(t *testing.T) {
	result := kidsWithCandies([]int{4, 2, 1, 1, 2}, 1)
	if !reflect.DeepEqual(result, []bool{true, false, false, false, false}) {
		t.Errorf("kids with candies should display [true, false, false, false, false] but got %v", result)
	}
}

func TestKidsWithCandiesShouldDisplay2TrueAnd1False(t *testing.T) {
	result := kidsWithCandies([]int{12, 1, 12}, 10)
	if !reflect.DeepEqual(result, []bool{true, false, true}) {
		t.Errorf("kids with candies should display [true, false, true] but got %v", result)
	}
}
