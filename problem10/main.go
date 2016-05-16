/* Problem 10
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million.
*/
package main

import (
    "fmt"
)


func seive(max int)(int){
    answer := 2
    primes := make([]int, 1)
    primes[0] = 2

    for i := 3; i < max; i++ { //i is a possible prime
        isPrime := true
        for j := 0; j < len(primes)-1; j++ {

            if i % primes[j] == 0 {
                isPrime = false
                break
            }
            if primes[j]*primes[j] > i  {
                break
            }
        }
        if isPrime {
            primes = append(primes, i)
            answer += i
        }
    }

    return answer
}

func main() {
    fmt.Println(seive(2000000))
}
