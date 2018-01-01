package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findAllDigitPowers(5))
}
func findAllDigitPowers(p int) (sum int) {
	for i := 2; i < 1000000; i++ {
		if i == calculateDigitPower(i,p) {
			fmt.Println(i)
			sum += i
		}
	}
	return
}
func calculateDigitPower(n,p int) (m int) {
	for n > 0 {
		t := n % 10
		m += int(math.Pow(float64(t),float64(p)))
		n /= 10
	}
	return
}

