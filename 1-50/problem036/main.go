package main

import (
	"fmt"
	"strconv"
	"time"
)
//find the sum of all numbers n < 1000000
//where n is a palindrome in base 10 and 2

func main() {
	t0 := time.Now()
	ans := problem36()
	fmt.Println("Answer: ",ans, "Time: ",time.Since(t0))
}
func problem36() (ans int) {
	for i := 0; i < 1000000; i++ {
		if isPalindrome(strconv.Itoa(i)) && isPalindrome(strconv.FormatInt(int64(i),2)) {
			ans += i
		}
	}
	return
}
func isPalindrome(n string) (bool) {
	for i := 0; i < (len(n) / 2); i++ {
		if n[i] != n[len(n) - i - 1] {
			return false
		}
	}
	return true
}