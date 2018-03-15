package main

import (
	"time"
	"math/big"
	"fmt"
)
//Problem:
//Find the largest x part of the fundamental solution to all 
//Pell Equations; n \in [2,1000].
//Research:
//https://en.wikipedia.org/wiki/Pell%27s_equation
//https://en.wikipedia.org/wiki/Continued_fraction
func main() {
	t0 := time.Now()
	ans := findMaximumXValue(1000)
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
}



func findMaximumXValue(n int) (index int){
	x := big.NewInt(0)
	for i := 2; i <= n; i++ {
		h := findMinimalSolution(i)
		if h == nil {
			continue
		}

		if h.Cmp(x) == 1{
			x = h
			index = i
		}
	}  
	return
}
//Finds minimal solution to Pell Equation of order n
func findMinimalSolution(n int) (*big.Int){
	root := big.NewFloat(float64(n))
	root.SetPrec(2000)
	root.Sqrt(root)
	floor := big.NewInt(0)
	h1,h2 := big.NewInt(1),big.NewInt(0)
	k1,k2 := big.NewInt(0),big.NewInt(1)
	//calculates the continued fraction representation of sqrt(n)
	//then calculates convergents until a pair of them satisfies h1*h1 - n*k1*k1 = 1
	for i := 0;; i++ {
		root.Int(floor) // a
		if root.IsInt() {
			return nil //perfect squares have no integer solutions
		}
		//new convergent h = a(n) * h(n-1)  + h(n-2)
		//same with k
		//h
		tmpH := big.NewInt(0)
		tmpH.Set(h1)
		h1.Mul(floor,tmpH)
		h1.Add(h1,h2)
		h2 = tmpH
		//k
		tmpK := big.NewInt(0)
		tmpK.Set(k1)
		k1.Mul(floor,tmpK)
		k1.Add(k1,k2)
		k2 = tmpK

		if isValidSolution(h1,k1,big.NewInt(int64(n))) {
			return h1
		}
		
		tmp := big.NewFloat(0)
		tmp.SetInt(floor)
		root.Sub(root,tmp)
		root.Quo(big.NewFloat(1),root)
	}
	return nil
}


//checks if a pair satisfies equation
func isValidSolution(h1,k1, n *big.Int) (bool) {
	h := big.NewInt(0)
	h.Set(h1)
	k := big.NewInt(0)
	k.Set(k1)
	val := big.NewInt(0)
	val.Sub(h.Mul(h,h), k.Mul(n,k.Mul(k,k)))
	return val.Cmp(big.NewInt(1)) == 0
	
}