package main

import (
	"fmt"
	"time"
)

const max = 1000 //Guess

func main() {
	t0 := time.Now()
	ans := findAnswer()
	fmt.Println("Answer:", ans, "\nTime:", time.Since(t0))
}

func findAnswer() int {
	cn := c()
	bn := b(cn)
	for i := range bn {
		if bn[i] > 5000 {
			return i
		}
	}
	return -1 // not found
}

//Black Magic Euler Transformation
//http://mathworld.wolfram.com/EulerTransform.html
func b(cn []int) []int {
	bn := make([]int, len(cn))
	bn[1] = cn[1]
	for i := 2; i < len(cn); i++ {
		bn[i] = cn[i]
		for k := 1; k < i; k++ {
			bn[i] += cn[k] * bn[i-k]
		}
		bn[i] /= i
	}
	return bn
}

//c[j] = sum of distinct prime factors of j, using modified seive of eratosthenes
func c() []int {
	cn := make([]int, max+1)
	an := make([]int, max+1)
	for i := range an {
		an[i] = 1
	}
	for i := 2; i < len(an); i++ {
		if an[i] == 1 {
			for j := i; j < len(an); j += i {
				cn[j] += i
			}
			for j := 2 * i; j <= max; j += i {
				an[j] = 0
			}
		}
	}
	return cn
}
