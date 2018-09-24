package main

import (
	"time"
	"fmt"
)

func main() {
	t0 := time.Now()
	ans := findAnswer(1000)
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
	// t1 := time.Now()
	// ans2 := findAnswer2(100)
	// fmt.Println("Answer 2:",ans2,"\nTime:",time.Since(t1))
	// for i := 1; i < 20; i++ {
	// 	//fmt.Print(i,":")
	// 	//Does number of solutions relate to prime factors or distinct prime factors?
	// 	//2^n seems to have n + 1 solutions
	// 	countSolns(i)
	// 	fmt.Println()
	// }
}

func findAnswer(target int) (int) {
	for i := 5; ; i++ {
		if i < 0 {
			fmt.Println("Overflow!")
			return -1
		}
		s := countSolns(i)
		if s > target {
			return i
		}
	}

}


func countSolns(n int) (s int) {
	for m := 1; m <= n; m++ {
		if n*n % m == 0 {
			s++
		}
	}
	return s
}