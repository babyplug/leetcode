package defang_ip

import (
	"testing"
)

func TestDefangIPSHouldDisplay8Of1(t *testing.T) {
	result := defangIPaddr("1.1.1.1")
	if result != "1[.]1[.]1[.]1" {
		t.Errorf("defangIPaddr should display '1[.]1[.]1[.]1' but got %v", result)
	}
}

func TestDefangIPSHouldDisplay255Dot100Dot50Dot0(t *testing.T) {
	result := defangIPaddr("255.100.50.0")
	if result != "255[.]100[.]50[.]0" {
		t.Errorf("defangIPaddr should display '255[.]100[.]50[.]0' but got %v", result)
	}
}
