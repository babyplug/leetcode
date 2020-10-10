package jewels_stones

import "testing"

func TestCase1ShouldDisplay3(t *testing.T) {
	result := numJewelsInStones("aA", "aAAbbbb")
	if result != 3 {
		t.Errorf("numJewelsInStone should display %v but result got %v \n", 3, result)
	}
}

func TestCase1ShouldDisplay0(t *testing.T) {
	result := numJewelsInStones("z", "ZZ")
	if result != 0 {
		t.Errorf("numJewelsInStone should display %v but result got %v \n", 0, result)
	}
}
