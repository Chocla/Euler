/*
Problem 4
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
*/
//extra: use concurrency somehow

package main

import (
    "fmt"

)

func product(min, max int)(int){
    answer := 0
    for i := min; i <= max; i++ {
        for j := min; j <= max; j++ {
            if palindrome(i*j) {
                if i*j > answer {
                    answer = i*j
                }
            }
        }
    }
    return answer
}
func palindrome(x int)( isPalindrome bool){
    isPalindrome = false
    original := x
    var tmp int
    for x > 0 {
        tmp *=10
        tmp += x%10
        x /=10
    }
    if original == tmp {
        isPalindrome = true
    }
    return isPalindrome
}
func main() {
    fmt.Println(product(100,999))
}
