package main

import (
	"fmt"
)
	//How many values of nCr with n \in [1,100] are greater than 1 million
	//note that the maximum r value is n/2
	//for each n 23 to 100
		//for each r from n/2 to 0 until nCr < 1000000
			//add 2 to the answer counter
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
		
		//fmt.Println(count)
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
