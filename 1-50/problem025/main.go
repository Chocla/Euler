package main


import (
	"fmt"
	"math/big"
)
/*What is the index of the first term in the 
 Fibonacci sequence to contain 1000 digits? */
func main() {
	a, b := big.NewInt(1), big.NewInt(1)
	var target big.Int
	target.Exp(big.NewInt(10),big.NewInt(999),nil)
	i := 1
	for a.Cmp(&target) < 0 {
		a.Add(a,b)
		a,b = b,a
		i++
	}
	fmt.Println(a)
	fmt.Println("Answer: ", i)
}
