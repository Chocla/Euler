package main

import "testing"
func TestCalc(t *testing.T) {
	i := 8208
	if i != calculateDigitPower(i,4) {
		t.Errorf("Uh oh")
	}
}