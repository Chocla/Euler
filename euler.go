/*
Project Euler Problems Written in Go (poorly)
*/

/* Problem 5
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
package main

import (
    "fmt"
)

func main() {
    answer := 0
    for i := 1;  ; i++ {
        for j := 1; j <= 20; j++ {
            if i % j == 0 && j == 20 { answer = i}
            if i % j == 0 {continue}
            if i % j != 0 {break}
        }
        if answer > 0 {
            break
        }
    }
    fmt.Println(answer)
}
