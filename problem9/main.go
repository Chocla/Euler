/* Problem 9
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a2 + b2 = c2

For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/
package main

import (
    "fmt"
    "math"
)

func main() {
   answer := 0
    var c float64

    for i := 4; i < 1000 ; i++ {
        for j := 3; j < i ; j++ {
            c = math.Sqrt(float64(i*i+j*j))
            if c - float64(int(c)) == 0 {

                if i+j+int(c) == 1000 {

                    answer = i*j*int(c)
                    fmt.Println(answer)
                    return
                }

            }
        }
    }

}
