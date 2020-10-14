package shuffle_array

import (
	"reflect"
	"testing"
)

func TestValueShouldDisplay2_3_5_4_1_7(t *testing.T) {
	result := shuffle([]int{2, 5, 1, 3, 4, 7}, 3)
	expect := []int{2, 3, 5, 4, 1, 7}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("shuffle should display [2, 3, 5, 4, 1, 7] but got %v", result)
	}
}
