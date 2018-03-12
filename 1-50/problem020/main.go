package main

import (
    "fmt"
    "os"
    "strconv"
    "math/big"
)

func factorial(n int, ) (*big.Int) {
    tmp := big.NewInt(int64(n))
    product := big.NewInt(1)

    for i := n; i > 1; i-- {
        product *=n
    }

}
func main() {
    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        panic(err)
    }
    prod := 1

    factorial(n)
    fmt.Println(prod)

}
