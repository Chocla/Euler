package main

import (
	"fmt"
	"time"
)


//find last 10 digits of 
//28433*(2**(7830457)) + 1
func main() {
	t0 := time.Now()
	ans := 1
	for i := 1; i <= 7830457; i++ {
		ans = (ans * 2 ) % 10000000000
	}
	ans = (ans *28433) % 10000000000
	ans += 1
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
}