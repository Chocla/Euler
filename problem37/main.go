package main

import (
	"fmt"
)
const primeMax = 1000000
func main() {
	sum := findAllTruncatablePrimes()
	fmt.Println(sum)
}

func findAllTruncatablePrimes() (sum int){
	pr := seive(primeMax)
	count := 0
	for _,p := range pr {
		if p < 10 {
			continue
		}
		if truncate(p,pr) {
			count++
			sum += p
		}
	}
	if count != 11 {
		fmt.Println("Expected 11 Primes. Found", count)
	}
	return
}

func truncate(p int, primes []int) (bool) {
	return truncateLeft(p, primes) && truncateRight(p,primes)
}
func truncateLeft(p int, primes []int) (bool) {
	for factor := 10; factor < p; factor *= 10 {
		flag := false
		tmp := p % factor
		for _,v := range primes {
			if tmp == v {
				flag = true
			}
		}
		if !flag {
			return false
		}
	}
	return true
}

func truncateRight(p int, primes []int) (bool) {
	for p > 10 {
		flag := false
		p /= 10
		
		for _,v := range primes {
			if p == v {
				flag = true
			}
		}
		if !flag {
			return false
		}
	}
	return true
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