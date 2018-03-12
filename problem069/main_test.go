package main

import "testing"

func TestPrimeFactors(t *testing.T) {
	tests := []struct {
		n int
		f []int
	}{
		{12, []int{2,3}},
		{315,[]int{3,5,7}},
	}
	for i := range tests{
		factors := primeFactors(tests[i].n)
		if !sliceEq(factors,tests[i].f) {
			t.Errorf("Oh No, Expected: %v, Got %v",tests[i],factors)
		}
	}
}
func TestTotient(t *testing.T) {
	tests := []struct {
		n int
		t int
	} {
		{2,1},
		{5,4},
		{7,6},
	}
	for i := range tests {
		T := totient2(tests[i].n)
		if T != tests[i].t {
			t.Errorf("Oh No! Expected: %v, Got: %v",tests[i],T)
		}
 	}
}
func sliceEq(a,b []int) (bool){
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}