package main

import (
	"fmt"
	"time"
)
//What is the smallest odd composite that cannot be
//written as the sum of a prime and twice a square?
const max = 10000
func main() {
	t0 := time.Now()
	ans := 0
	primes := seive(max)
	
	for i := 9; i < max;i += 2 {
		found := false
		if isPrime(i,primes) {
			continue
		} else {
			for _, p := range primes {
				if p > i -2 || found{
					break
				}
				for j := 1; 2*j*j < i-3 && !found; j++ {
					if i == p + 2*j*j {
						
						found = true
					}
				}
			}
		}
		if !found {
			ans = i
			break
		}
	}
	fmt.Println("Answer: ", ans, "Time: ", time.Since(t0))
}
func isPrime(n int, primes []int) (bool) {
	for i := range primes {
		if primes[i] == n {
			return true
		}
	}
	return false
}
func seive(upperBound int) (primes []int) {
	nums := make([]bool,upperBound)
	for i := 2; i < upperBound;i++ {
		if !nums[i] {
			primes = append(primes,i)
			//nums[i] = true
			for j := 2*i; j < upperBound; j+= i {
				nums[j] = true
			}
		}
	}

	return
}