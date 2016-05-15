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

func genproduct(c chan int64, x int64) {
    for i := 100; i < 1000; i++ {
        for j := 999; j <= 1000-i; j++ {
            c <- int64(i*j)
        }
    }
}
func palindrome(c chan int64)(bool, int64){
    num := <- c
    var newnum int
    var isPalindrome = false
    for num > 0 {
        newnum *=10
        newnum += num%10
        num /=10
    }
    if num == newnum {
        isPalindrome = true
    }
    return isPalindrome, num
}
func main() {
    c := make(chan int64)
    go genproduct(c,)


}
