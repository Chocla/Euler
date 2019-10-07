package main

import (
	"time"
	"fmt"
)
//factorial map
var fact = map[int]int{
	0:1,
	1:1,
	2:2,
	3:6,
	4:24,
	5:120,
	6:720,
	7:5040,
	8:40320,
	9:362880,
}
//map of loop values
var ends = map[int]int {
	1:1,
	2:1,
	145:1,
	169:3,
	363601:3,
	1454:3,
	871:2,
	45361:2,
	872:2,
	45362:2,
	40585:1,
}
//if a number was in a previously computed chain, we don't need to calculate it
var seenNums = make([]bool,10000001)
func main() {
	t0 := time.Now()
	ans := findAnswer()
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
}

func findAnswer() (int){
	answer := 0
	for i := 2; i < 1000000; i++ {
		len := computeChainLength(i)
		if len == 60 {
			answer++
		}
	}
	return answer
}

//compute the length of unique chain of numbers 
//where each number is the digit factorial sum of the previous
func computeChainLength(n int) (length int){
	length = 1
	if !seenNums[n] {
	for {
		n = digitFact(digitize(n))
		if ends[n] != 0 {
			length += ends[n]
			return 
		}

		seenNums[n] = true
		if length == 60 {
			return 
		}
		length++
	}
	}
	return 0
}

//sum factorials of digits
func digitFact(d []int) (n int){
	for _,v := range(d) {
		n += fact[v]
	}
	return
}
//turn integer into slice of digits
func digitize(n int) (d []int) {
	for n > 0 {
		d = append(d,n%10)
		n /= 10
	}
	return
}