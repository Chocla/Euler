package main

import (
	"time"
	"fmt"
	"math/big"
)

func main() {
	t0 := time.Now()
	ans := findAnswer()
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
}

func findAnswer() (ans int){
	for i := 1; i <= 1000; i++ {
		approx := generateApproximation(i)
		if hasProperty(approx) {
			ans++
		}
	}	
	return
}
//generates the continued fraction approximation of sqrt(2) n layers deep, simplified
func generateApproximation(n int) ( *big.Rat){
	two := big.NewRat(2,1)
	a := big.NewRat(1,2)
	for i := 1; i < n; i++ {
		a.Add(a,two)
		a.Inv(a)
	}
	a.Add(a,big.NewRat(1,1))
	return a
}
//checks if the numerator has more digits than the denominator
func hasProperty(x *big.Rat) (bool) {
	num := x.Num().String()
	den := x.Denom().String()
	if len(num) > len(den) {
		return true
	}
	return false
}