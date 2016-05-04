/*
Project Euler Problems Written in Go (poorly)
*/
/*
Problem 7
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10 001st prime number?
*/

package main

import (
    "fmt"

)

func seive(target int)( int){
    primes := make( []int,1)
    primes[0] = 2
    primeCount := 1
    for candidate := 3; ; candidate++ {
        var isPrime bool = true
         for i := 0; i < primeCount; i++ {
             test := primes[i]
             //fmt.Println(test, isPrime)
             if candidate % test == 0 {
                 isPrime = false
                 break;
             }
         }
         if isPrime {
            // fmt.Println(primes)
             primes = append(primes, candidate)
             primeCount++
             if primeCount == target {
                 return primes[target-1]
             }
         }
    }
}
func main() {

    fmt.Println(seive(10001))
}
