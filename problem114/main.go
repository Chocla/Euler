package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	ans := memoizedBlockCombinations(50)
	fmt.Println("Answer:", ans, "\nTime:",time.Since(t0))
}

func memoizedBlockCombinations(n int) (int) {
	n += 1
	c := make([]int, n+ 1)
	for i := 0; i <= n; i++ {
		c[i] = 1
		// initial block size
		for k := 3; k <= n; k++ {
			for j := 0; j <= n-k; j++ {
				if i-j-k-1 >= 0 {
					c[i] += c[i-j-k-1]
				}
			}
			
		}
	}
	return c[len(c) - 1]
}