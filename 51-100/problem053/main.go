package main

import (
	"fmt"
)

func main() {
	fmt.Println(findAnswer())
}

func findAnswer() (count int){
	for n := uint64(23); n <= 100; n++ {
		r := uint64(1)
		for ; r <= n/2; r++ {
			if comb(n,r) > 1000000 {
				
				break
			}
		}
			count += int(n - 2*r +1)
		
	}
	return 
}
func comb(n, r uint64) (ans uint64) {
	ans = 1
	if r > n {
		return 0
	}
	for i := uint64(1); i <= r; i++ {
		ans *= n
		n--
		ans /= i
		//fmt.Println(ans, i, n)
	}
	return
}
