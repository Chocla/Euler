package main

import (
	"time"
	"math/big"
	"fmt"
)

func main() {
	t0 := time.Now()
	ans := findOddPeriodsBelow(10000)
	fmt.Println("Answer: ", ans, "\nTime: ",time.Since(t0))
}

func findOddPeriodsBelow(uBound int) (count int){
	for i := 2; i <= uBound; i++ {
		a := findPeriod(i)
		if  a % 2 != 0 { //odd period
			count++
		}
	}
	return
}
//caluculates the period of a sqrt's continued fraction expansion
func findPeriod(val int) (period int) {
	sequence := make([]int64,0)
	root := big.NewFloat(float64(val))
	root.SetPrec(3000) //set precision very high so that long sequences don't get messed up from rounding
	root.Sqrt(root)
	floor := big.NewInt(0)

	//algorithm for computing the continued fraction terms
	for i := 0;; i++ {
		root.Int(floor) 
		if root.IsInt() {
			return 0 //ignores perfect squares
		} else {
			sequence = append(sequence, floor.Int64())
		}
		f,a := hasPattern(sequence[1:])
		
		if a == -1 {
			return 1
		}
		if f {
			return (len(sequence) - 1 )/4
		}
		tmp := big.NewFloat(0)
		tmp.SetInt(floor)
		root.Sub(root,tmp)
		root.Quo(big.NewFloat(1),root)
	}
}

func hasPattern(seq []int64) (bool,int) {
	if len(seq) % 4 != 0 || len(seq) == 0 {
		return false,0
	}

	//hack-y fix for fractions with period 1
	allOneNum := true
	num := seq[0]
	for i := 0; i < len(seq); i++ {
		if seq[i] != num {
			allOneNum = false
			break
		}
	}
	if allOneNum {
		if len(seq) > 9 {
			return true,-1
		}
		return false,0
	}

	//checks if a pattern repeats 4 times before accepting that it is a pattern
	for i := 0; i < len(seq)/4; i++ {
		if seq[i] != seq[i +   len(seq)/4] || 
		   seq[i] != seq[i +   len(seq)/2] || 
		   seq[i] != seq[i + 3*len(seq)/4] {
			return false, 0
		}
	}

	return true,0
}
