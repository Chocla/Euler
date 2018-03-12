package main

import "testing"
import "strconv"

func TestIsPalindrome(t *testing.T) {
	a,b := strconv.FormatInt(585,2), "544"
	if !isPalindrome(a) {
		t.Errorf("uh oh: %v",a)
	}
	if isPalindrome(b) {
		t.Errorf("uh oh: %v",b)
	}

} 