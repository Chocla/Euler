package main

import (
	"fmt"
	"time"
)
const size = 12000
func main() {
	t0 := time.Now()
	ans := findAnswer()
	fmt.Println("Answer:", ans, "\nTime:", time.Since(t0))
}
//the number of fractions between 1/3 and 1/2 is the distance between their indices
func findAnswer() (int) {
	return (findIndex(size,1,2) - findIndex(size,1,3) - 1)
}

//find's the index of desired fraction in F(N)
//N: order of Farey Sequence
//n: numerator of desired fraction
//d: denominator of desired fraction
func findIndex(N, num, den int) (int) {
	i := 0
	a, b, c, d := 0,1,1,N
	for (a != 1 || b != 1){
		//the next term is defined in terms of the previous two terms:
		//p = floor((N + b) / d) * c - a
		//q = floor((N + b) / d) * d - b
		k := (N + b) / d
		a, b, c, d = c, d, (k * c - a), (k * d - b)
		if a == num && b == den {
			return i
		}
		i++
	}
	return i
}