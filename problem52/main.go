package main

import (
	"time"
	"sort"
	"fmt"
)

//Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits.
func main() {
	t0 := time.Now()
	ans := findSolution()
	fmt.Println("Answer: ", ans, "\nTime: ",time.Since(t0))
}

//infinite loop to check for desired property
func findSolution() (ans int) {
	for i := 1; ; i++ {
		if checkProperty(i) {
			return i
		}
	}
	
}
//checks to see if the integer x has the desired property for the problem
func checkProperty(x int) (bool) {
	digits := digitify(x)
	sort.Ints(digits)
	for i := 2; i <= 6; i++ {
		newDigits := digitify(x*i)
		sort.Ints(newDigits)
		
		if !sliceEq(newDigits,digits){
			return false
		}
	} 
	return true
}

//turns an integer x into a slice containing its digits
func digitify(x int)(digits []int) {
	for ;x > 0; x /= 10 {
		digits = append(digits, x % 10) 
	}
	return 
}

//check if two slices are equal
func sliceEq(a,b []int) (bool){
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}