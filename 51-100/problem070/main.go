package main

import (
	"time"
	"sort"
	"fmt"
)

func main() {
	t0 := time.Now()
	ans := findAnswer(10000000) 
	fmt.Println("Answer:", ans, "\nTime: ",time.Since(t0))

}

func findAnswer(max int) (int) {
	min := 10.0 //arbitrary val
	minIndex := 0
	phi := totientSeive(max) //populate totient values
	for i := 2; i < len(phi); i++ {
		if isPermutation(digitify(i), digitify(phi[i])) {
			//find n value for which n/φ(n) is minimum
			candidate := float64(i) / float64(phi[i]) 
			if  candidate < min {
				min = candidate
				minIndex = i
			}
		}
	}
	return minIndex

}

func totientSeive(max int) ([]int){
	phiVals := make([]int, max+1)
	for i := 0; i <= max; i++ {
		phiVals[i] = i
	}
	for i := 2; i < len(phiVals); i++ {
		//i is prime
		if phiVals[i] == i {
			phiVals[i] = i-1
			for j := 2*i; j < len(phiVals);j += i {
				phiVals[j] = (phiVals[j] / i) * (i-1) 
			}
		}
	}
	return phiVals
}

func digitify(n int) (d []int) {
	for n > 0 {
		d = append(d,n % 10)
		n /= 10
	}
	return
}
//if two numbers are permutations of each other, their sorted digits should be equal
func isPermutation(d1, d2 []int) (bool){
	//don't need to waste time sorting if they have different lengths
	if len(d1) != len(d2) {
		return false
	}

 	sort.Ints(d1)
	sort.Ints(d2)
	for i := range d1 {
		if d1[i] != d2[i] {
			return false
		}
	}

	return true
}