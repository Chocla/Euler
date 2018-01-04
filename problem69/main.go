package main

import (
	"fmt"
	"time"
)

//Find the value of n ≤ 1,000,000 
//for which n/φ(n) is a maximum.
func main() {
	t0 := time.Now()
	phiVals := initPhiVals(1000001)
	totientSeive(phiVals)
	ans := float64(0)
	ansIndex := 0
	for i := 2; i < len(phiVals);i++ {
		if ans < float64(i) / float64(phiVals[i]) {
			ans = float64(i) / float64(phiVals[i]) 
			ansIndex = i
		}
	}
	//Runs in < 1s
	fmt.Println("Answer: ",ansIndex,"Time: ",time.Since(t0))
}
//
func initPhiVals(n int) (phiVals []int) {

	for i := 0; i < n; i++ {
		phiVals = append(phiVals,i)
	}
	return
}
func totientSeive(phiVals []int) {
	for i := 2; i < len(phiVals); i++ {
		//i is prime
		if phiVals[i] == i {
			phiVals[i] = i-1
			for j := 2*i; j < len(phiVals);j += i {
				phiVals[j] = (phiVals[j] / i) * (i-1) 
			}
		}
	}
}