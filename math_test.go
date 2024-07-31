package main

import "testing"

func TestSoma1(t *testing.T) {
	total := soma(15, 15)	

	if (total != 30) {
		t.Errorf("Invalid sum result: %d. Expected: %d", total, 30)
	}
}

func TestSoma2(t *testing.T) {
	total := soma(2, 3)	

	if (total != 5) {
		t.Errorf("Invalid sum result: %d. Expected: %d", total, 5)
	}
}