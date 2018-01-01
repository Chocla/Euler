package main

import (
	"fmt"
)

func main() {
	a,b,maxL := problem27(-1000,1000)
	fmt.Printf("n^2 + %dn + %d Produces %d consecutive primes\n Answer: %d",a,b,maxL,a*b)
}

func problem27(min,max int) (a,b,maxL int) {

	var i,j int
	for i = min+1 ; i < max; i++ {
		for j = min ; j <= max ; j++ {
			t := primeCheck(i,j) 
			if t > maxL {
				maxL = t
				a,b = i,j
			}
		}
	}

	return a,b, maxL
}
func primeCheck(a,b int) (int) {
	i := 0
	for ; isPrime(i*i + a*i + b) ;i++ {

	}

	return i - 1
}

func isPrime(n int) (bool) {
	if n < 1 {
		return false
	}
	for i := 2; i*i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}