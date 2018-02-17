package main

import (
	"time"
	"fmt"
	"math/big"
)
//Considering natural numbers of the form, a^b, where a, b < 100, what is the maximum digital sum?

func main() {
	t0 := time.Now()
	ans := findMaximumDigitalSum()
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
}
func findMaximumDigitalSum() (max int) {
	for a := 1; a < 100; a++ {
		for b := 1; b < 100; b++ {
			candidate := sumDigits(generatePower(a,b))
			if candidate > max {
				max = candidate
			}
		}
	}
	return
}
//create string with value a^b
func generatePower(a,b int) (string){
	x := big.NewInt(int64(a))
	x.Exp(x, big.NewInt(int64(b)),nil)
	return x.String()
}

func sumDigits(num string) (sum int){
	for i := range num {
		sum += int(num[i] - '0') //convert ascii representation of a number into actual value ('1' - '0' = 1)
	}
	return 
}