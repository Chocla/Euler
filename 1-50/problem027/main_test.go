package main

import "testing"

func TestIsPrime(t *testing.T) {
	a,b := 40, 41
	if isPrime(a) {
		t.Errorf("IsPrime Failure: %d is prime, expected composite",a)
	}
	if !isPrime(b) {
		t.Errorf("IsPrime Failure: %d is composite, expected prime",b)
	}
}

func TestPrimeCheck(t *testing.T) {
	a,b := 1,41
	c := primeCheck(a,b)
	if c != 40 {
		t.Errorf("Error: Expected 40, got %d",c)
	}
	a,b = -79,1601
	c = primeCheck(a,b)
	if c != 80 {
		t.Errorf("Error: Expected 80, got %d",c)
	}
}