package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	ans := findAnswer(1000000)
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
}

//calculates the first 7 digits of p(n), until p(n) % 1000000 = 0
func findAnswer(target int) (int){
	part := []int{1}
	pents := []int{0}
	j := 1
	
	for i := 1; ; i++ {
		//We only need pentagonal numbers up to the size
		//of p(n) we want to calculate.
		if i + 1> pents[len(pents) - 1] {
			pents = append(pents,nextPent(j)...)
			j++
		}

		part = append(part, nextPart(part,pents))
		//guarantees value will be in interval [0,1000000)
		if part[i] < 0 {
			part[i] += 1000000

		}
		//answer found
		if part[i] % target == 0 {
			return i
		}
	}
}
//calcluates p(n) using Euler's Pentagonal Number Theorem
func nextPart(prevParts []int, pents []int) (pn int) {
	sign := 1
	n := len(prevParts)
	for i := 1; (n >= pents[i]) ; i++ {
		term := prevParts[n - pents[i] ] % 1000000
		pn += sign*term

		if i % 2 == 0 && i != 0{
			sign *= -1
		}
	}
	return pn % 1000000
}
//calculates the next two generalized pentagonal numbers (using n and -n)
func nextPent(n int) ([]int) {
	return []int{(3*n*n -n)/2, (3*n*n + n)/2 }
}