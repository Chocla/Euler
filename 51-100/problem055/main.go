package main

import (
	"time"
	//"strconv"
	"math/big"
	"fmt"
)

func main() {
	t0 := time.Now()
	ans := findLychrelNums(10000)
	fmt.Println("Answer:", ans, "\nTime:",time.Since(t0))
}

func findLychrelNums(uBound int) (ans int){
	for i := 0; i <= uBound; i++ {
		if isLychrel(i) {
			ans++
		}
	}
	return
}

func isLychrel(x int) (bool) {
	bigX := big.NewInt(int64(x))
	for i := 0; i <= 50; i++ {
		r := big.NewInt(int64(0))
		r.SetString(rev(bigX.String()),10)
		bigX.Add(bigX,r)
		if isPalindrome(bigX.String()) {
			return false
		}
		
	}
	return true
}

func rev(s string) (r string) {
	for i := len(s)-1; i >= 0; i--{
		r += string(s[i])
	}
	return
}
func isPalindrome(x string) (bool) {
	for i := 0; i < len(x) / 2; i++ {
		if x[i] != x[len(x)-i-1] {
			return false
		}
	}
	return true
}