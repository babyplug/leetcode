package sort

import (
	"reflect"
	"testing"
)

func TestShoudSortAscending(t *testing.T) {
	testCase := []int{4, 3, 2, 10, 12, 1, 5, 6}

	bubbleSorted := bubbleSort(testCase)
	selectionSorted := selectionSort(testCase)
	insertionSorted := insertionSort(testCase)

	if !reflect.DeepEqual(bubbleSorted, []int{1, 2, 3, 4, 5, 6, 10, 12}) {
		t.Errorf("bubbleSort should sort ascending but got %v", bubbleSorted)
	}

	if !reflect.DeepEqual(selectionSorted, []int{1, 2, 3, 4, 5, 6, 10, 12}) {
		t.Errorf("selectionSort should sort ascending but got %v", selectionSorted)
	}

	if !reflect.DeepEqual(insertionSorted, []int{1, 2, 3, 4, 5, 6, 10, 12}) {
		t.Errorf("insertionSort should sort ascending but got %v", insertionSorted)
	}

}
