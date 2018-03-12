package main

import (
	"fmt"
	"math/big"
)
//How many n-digit positive integers exist which are also an nth power?
func main() {
	
	fmt.Println(aa())
}

func aa() (ans int) {
	ans = 1
	for i := 2;i <= 10; i++ {
		for j := 1; ; j++ {
			pow := generatePower(i,j)
			if countDigits(pow) == j {
				ans++
				fmt.Println(pow)
			}
			if countDigits(pow) > j || i*i < j{
				break
			}

		}
	}
	return
}
func generatePower(a,b int) (string){
	x := big.NewInt(int64(a))
	x.Exp(x, big.NewInt(int64(b)),nil)
	return x.String()
}
func countDigits(num string) (int) {
	return len(num)
}