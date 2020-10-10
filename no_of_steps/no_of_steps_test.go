package no_of_steps

import (
	"fmt"
	"testing"
)

func TestCaseVal14ShouldDisplay6(t *testing.T) {
	result := numberOfSteps(14)
	fmt.Println(result)
	if result != 6 {
		t.Errorf("numberOfSteps result should display %v but bot %v", 6, result)
	}
}

func TestCaseVal8ShouldDisplay4(t *testing.T) {
	result := numberOfSteps(8)
	if result != 4 {
		t.Errorf("numberOfSteps result should display %v but bot %v", 4, result)
	}
}

func TestCaseVal123ShouldDisplay12(t *testing.T) {
	result := numberOfSteps(123)
	if result != 12 {
		t.Errorf("numberOfSteps result should display %v but bot %v", 12, result)
	}
}
