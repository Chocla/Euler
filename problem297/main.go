package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	ans := helperZeck(100000000000000000)
	fmt.Println(ans,time.Since(t0))

}

func helperZeck(max int) (int) {
	fib := firstNFib(84)
	sfib := secondOrderFib(84,fib)
	return recursiveZeck(max,fib,sfib)
}

func recursiveZeck(max int, fib []int, sfib []int) (sz int){
	if max == 0 {
		return 
	} else {
		i := 1
		for ; fib[i+1] < max; i++ {
			sz += sfib[i-1]
		}
		return sz + (max - fib[i]) + recursiveZeck(max - fib[i], fib, sfib)
	}
}

//return first n fibonacci numbers
func fib(n int) ([]int) {
	fib := make([]int, n+1)
	fib[1] = 1
	for i := 2; i <=n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}
//return first n second order fibonacci numbers
func secondOrderFib(n int, fib []int) ([]int) {
	sfib := make([]int,n+1) 
	sfib[1] = 1
	sfib[2] = 1
	for i := 3; i <= n; i++ {
		sfib[i] = sfib[i-1] + sfib[i-2] + fib[i-2]
	}

	return sfib
}