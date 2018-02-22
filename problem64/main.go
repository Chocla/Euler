package main

import (
	//"math"
	"math/big"
	"fmt"
)

func main() {
	//fmt.Println(findPeriod(9904))
	fmt.Println(findOddPeriodsBelow(10000))
}
func findOddPeriodsBelow(uBound int) (count int){
	
	for i := 2; i <= uBound; i++ {
		a := findPeriod(i)
		fmt.Println(i,a)
		if a == -1 {
			fmt.Println("leak for i = ", i)
			panic("uh oh")
		}
		if  a % 2 != 0 {
			count++
		}
	}
	return
}
//caluculates the period of a sqrt's continued fraction expansion

func findPeriod(val int) (period int) {
	sequence := make([]int64,0)
	root := big.NewFloat(float64(val))
	root.SetPrec(3000)
	root.Sqrt(root)
	floor := big.NewInt(0)
	a := 0
	secondaryFlag := false
	f := false
	for i := 0;; i++ {
		root.Int(floor) //take floor
		if root.IsInt() {
			return 0 
		} else {
			sequence = append(sequence, floor.Int64())
		}
		//fmt.Println(a, floor.Int64())
		if secondaryFlag && int64(a) == floor.Int64(){
			return (len(sequence) - 1 )/2
		} else {
			secondaryFlag = false
		}
		f,a = hasPattern(sequence[1:])
		
		if a == -1 {
			return 1
		}

		if f {
			secondaryFlag = true
			//return (len(sequence) - 1 )/2
		}
		tmp := big.NewFloat(0)
		tmp.SetInt(floor)
		//fmt.Println(root.String(),floor.String(),a)
		root.Sub(root,tmp)
		root.Quo(big.NewFloat(1),root)
	}
		
	
	//ENDFOR

	return
}
//improve to require pattern to occur 3 times before counting it
func hasPattern(seq []int64) (bool, int) {
	if len(seq) % 2 != 0 || len(seq) == 0{
		return false, 0
	}
	allOneNum := false
	num := seq[0]
	for i := 0; i < len(seq); i++ {
		if seq[i] != num {
			break
		}
		if i == len(seq) - 1 {
			allOneNum = true
		}
	}
	if allOneNum {
		if len(seq) > 9 {
			return true,-1
		}
		return false,0
	}
	for i := 0; i < len(seq) / 2; i++ {
		if seq[i] != seq[i + (len(seq)/2)] {
			return false,0
		}
	}
	return true,int(seq[0])
}