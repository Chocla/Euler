package main

import (
    "fmt"
    "os"
    "strconv"
)

func countDivisor(n int)(int)  {
    divisors := 0
    for i := 1; i*i < n; i++ {
        if n % i == 0 {
            divisors+=2
        }
    }
    return divisors
}

func main()  {
    if len(os.Args) != 2{
        panic("enter a number")
    }
    target,_ := strconv.Atoi(os.Args[1])
    triangle := 0
    for i := 1; ; i++ {
        triangle += i
        if countDivisor(triangle) >= target {
            fmt.Println(triangle)
            break
        }

    }
}
