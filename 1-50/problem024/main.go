package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	test := []int{0,1,2,3,4,5,6,7,8,9}
	
	for i := 2; i <= 1000000; i++ {
		test = nextPerm(test)
		//fmt.Println(test)
	}
	t1 := time.Now()
	fmt.Println(test, "in ", t1.Sub(t0))
}

func nextPerm(a []int) ([]int) {
	i := len(a) - 1
	for i > 0 && a[i-1] >= a[i] {
		i--
	}
	if i <= 0 {
		return a
	}
	j := len(a) - 1
	for a[j] <= a[i - 1] {
		j--
	}
	a[i-1],a[j] = a[j],a[i-1]
	
	j = len(a) - 1
	for i < j {
		a[i],a[j] = a[j],a[i]
		i++
		j--
	}
	return a
}
