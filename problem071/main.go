package main

import (
	"time"
	"math/big"
	"fmt"
)
const uBound = 1000000

func main() {
	t0 := time.Now()
	ans := findAnswer(uBound)
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
}
//find's the closest rational to 3/7 such that the denominator is < 1,000,000 and the gcd(num,deonm) = 1
func findAnswer(max int64) (int64) {
	lower := big.NewRat(2,5) //given lower bound
	upper := big.NewRat(3,7)
	flag := 0
	for d := max; flag < 10; d-- { 
	    for n := d / 2;  ; n-- {
			if gcd(n,d) == 1 {
				candidate := big.NewRat(n, d)
				//find the first n such that n/d is between the current lower bound and 3/7
				//reset the flag counter
				if candidate.Cmp(upper) == -1 && candidate.Cmp(lower) == 1{
					flag = 0
					lower = candidate
					break
				}
				//desired fraction doesn't exist for given d
				//if 10 denominators in a row don't work, assume we have found the closest num
				if candidate.Cmp(lower) == -1 {
					flag++
					break
				}
			}
		}
		
	}
	return lower.Num().Int64()
}
//simple GCD algorithm
func gcd(n, d int64) (g int64){
	for d != 0 {
		g = n % d
		n = d
		d = g
	}
	return n
}