package main

import (
	"strconv"
	"math"
	"time"
	"fmt"
)
/**
Some positive integers n have the property that the sum 
[ n + reverse(n) ] consists entirely of odd 
(decimal) digits. For instance, 36 + 63 = 99 and
409 + 904 = 1313. We will call such numbers reversible;
so 36, 63, 409, and 904 are reversible.
Leading zeroes are not allowed in either n or reverse(n).

There are 120 reversible numbers below one-thousand.

How many reversible numbers are there below one-billion?
*/
const max = 100000000
//const max = 1000
func main() {
	//brute force solution ~2 minutes
	t0 := time.Now()
	count := 0
	for i := 1; i < max; i++ {
		// if i % 1000000 == 0 {
		// 	fmt.Println(i, time.Since(t0),count)
		// }
		r,f := reverse(i)
		if f {
			if checkProperty(i + r) {
				count++
			}
		}
	}
	fmt.Println("Answer: ",count, "Time:", time.Since(t0))
}
func checkProperty(n int) (bool) {
	for n > 0 {
		if (n % 10) % 2 == 0 {
			return false
		}
		n /= 10
	}
	return true
}
func reverse(n int) (r int, flag bool) {
	flag = true
	l := len(strconv.Itoa(n))
	i := math.Pow(10, float64(l)-1)
	for ; n > 0;  {
		r += int(i) * (n % 10)
		n /= 10
		i /= 10
	}
	if len(strconv.Itoa(r)) != l {
		flag = false
	}
	return
}