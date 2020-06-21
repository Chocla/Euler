package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	ans := findAnswer(50)
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
}

func findAnswer(m int) int {
	for i := 10; ; i++ {
		// fmt.Println(i,memoizedBlockCombinations(50,i))
		if memoizedBlockCombinations(50,i) > 1000000 {
			return i
		}
	}

}

func memoizedBlockCombinations(m,n int) (int) {
	n += 1
	c := make([]int, n+ 1)
	for i := 0; i <= n; i++ {
		c[i] = 1
		// initial block size
		for k := m; k <= n; k++ {
			for j := 0; j <= n-k; j++ {
				if i-j-k-1 >= 0 {
					c[i] += c[i-j-k-1]
				}
			}
			
		}
	}
	return c[len(c) - 1]
}