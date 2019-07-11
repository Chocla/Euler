package main

import (
	"fmt"
	"math"
	"time"
)

const (
	MAX = 333333333
)
func main() {
	findAnswer(MAX)	
}

func findAnswer(max int64) {
	t0 := time.Now()
	ans := int64(0)
	for i := int64(2); i <= max; i++ {
		ans += sumIntegralTrianglePerimeters(i)
	}
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
} 

func sumIntegralTrianglePerimeters(a int64) (c int64) {
	n1,f1 := isPerfectSquare(3*a*a - 2*a - 1)
	n2,f2 := isPerfectSquare(3*a*a + 2*a - 1)

	if f1 && n1*(a+1) % 4 == 0{
		c += 3*a+1
	}
	if f2 && n2*(a-1) % 4 == 0{
		c += 3*a-1
	}

	return
}

func isPerfectSquare(n int64) (int64,bool) {
	tmp := int64(math.Sqrt(float64(n)))
	if tmp*tmp == n {
		return tmp,true
	}
	return 0,false
}