package main

import (
	"time"
	"fmt"
)

func main() {
	t0 := time.Now()
	ans := findSquare()
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
}
func findSquare() (int) {
	diags := make([]int, 0) 
	r := float64(0)
	for i := 3; ;i += 2 {
		diags = genDiags(i)
		r = calculateRatio(r,1 + 4*(i-1)/2,diags)
		if r < 0.1 {
			return i
		}
	}
	return 0
}

//calculates new ratio based on the number of primes in the new diagonals and the previous ratio
func calculateRatio(prevRat float64, length int, diags []int) (primeRat float64) {
	pCount := 0
	for i := range diags {
		if isPrime(diags[i]) {
			pCount++
		}
	}
	return ((prevRat*float64((length-4))) + float64(pCount)) / float64(length)
}

func genDiags(n int) ([]int){
	return []int{n * n, 
				  n * n -   (n - 1),
		          n * n - 2*(n - 1), 
				  n * n - 3*(n - 1)}
}
func isPrime(n int) (bool){
	for i := 2; i*i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}