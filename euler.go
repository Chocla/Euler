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
    "math"
)

func sumsquare(x int) (int){
    var y float64 = 0
    for i := 1; i <= x; i++ {
        y += math.Pow(float64(i),2)
    }
    return int(y)
}
func squaresum(x int) (int){
    y := 0
    for i := 1; i <= x; i++ {
        y += i
    }
    return int(math.Pow(float64(y),2))
}
func main() {
    answer := squaresum(100) - sumsquare(100)
    fmt.Println(answer)
}
