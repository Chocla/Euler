package main

import (
	"fmt"
	"strconv"
	"time"
)

//find the sum of palindromic numbers < 10^8 that can be written as a sum of consecutive integers
func main() {
	t0 := time.Now()
	n,s,p := computeConsecutiveSquares(100000000)
	s = findSum(p,s)
	fmt.Println("Num:", n, "Sum: ", s,"Time: ",time.Since(t0))
}

//correct for double counted palindromes
func findSum(p []int, s int) (int) {
	b := make([]bool,len(p))
	for i := range p {
		for j := range p {
			if p[i] == p[j] && i !=j && !b[i] && !b[j]{
				s -= p[i]
				b[i] = true
				b[j] = true
			}
		}
		b[i] = true
	}
	return s
}
//check all possible sequences of consecutive squares, and add the palindromes to sum
func computeConsecutiveSquares(upperBound int) (numPalindromes,sum int, pals []int) {
	for i := 1; i*i <= upperBound; i++ {
		tmp := i*i
		for j := i+1; j*j <= upperBound; j++ {
			tmp += j*j
			if tmp > upperBound {
				break
			}
			if isPalindrome(tmp) {
				//fmt.Println(tmp, i, j)
				sum += tmp
				numPalindromes++
				pals = append(pals,tmp)
			}
		}
	}

	return
}

func isPalindrome(n int) (bool) {
	str := strconv.Itoa(n)
	for i := 0; i < len(str) / 2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}