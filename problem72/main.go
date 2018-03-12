package main

import (
	"time"
	"fmt"
)

//problem is asking for the length of the Farey Sequence for n = 10^6
//|F(n)| = 1 + sum of totients of [1,n] - 2
//subtract two because problem doesn't count 0 or 1 in the sequence
func main() {
	t0 := time.Now()
	ans := sumTotients(totientSeive(1000000)) - 1
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
}

func totientSeive(n int) ( []int) {
	phiVals := initPhiVals(n)
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
func initPhiVals(n int) (phiVals []int) {

	for i := 0; i <= n; i++ {
		phiVals = append(phiVals,i)
	}
	return
}
func sumTotients(phiVals []int) (sum int) {
	for i := range phiVals {
		sum += phiVals[i]
	}
	return
}