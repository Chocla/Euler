package main

import (
	"time"
	"math/big"
	"fmt"
)

func main() {
	t0 := time.Now()
	ans := findNumeratorSum(100)
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
}

//calculates the nth term in the sequence of fraction approximations of e
//returns the sum of the digits in the numerator
func findNumeratorSum(n int) (ans int) {
	s := generateSequence(n)
	frac := makeFraction(s).Num().String()
	for i := range frac {
		ans += int(frac[i] - '0') //hack to calculate the integer value of a numeric character without converting it
	}

	return
}

//given a coefficient sequence, simplifies a finite continued fraction
func makeFraction(seq []int) (*big.Rat) {
	frac := big.NewRat(1,int64(seq[len(seq)-1]))
	for i := len(seq) - 2; i > 0; i-- {
		frac.Add(big.NewRat(int64(seq[i]),1),frac) //add the next term
		frac.Inv(frac) //find 1 / sum
	}
	frac.Add(big.NewRat(int64(seq[0]),1),frac) //add final term
	return frac
}
//generates the first n terms of coefficient sequence for the continued fraction representation of e
//[2,1,1,2,1,1,4,1,1,2k...]
func generateSequence(n int) ([]int) {
	if n < 0 {
		return nil
	}
	seq := make([]int, n)
	seq[0]=2
	k := 1
	for i := 1; i < len(seq);i++ {
		if i % 3 == 2 {
			seq[i] = 2*k
			k++
		} else {
			seq[i] = 1
		}
	}
	return seq
}