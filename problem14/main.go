package main

import (
    "fmt"
    "sync"
)

const max int = 1000000
type answer struct {
    value int
    length int
}

func collatz(n int,wg *sync.WaitGroup, length *int) {
    *length++
    if n == 1 {
        return
    }

    wg.Add(1)
    go func() {
        if n % 2 == 0 {
            collatz(n/2,wg,length)
        } else {
            collatz(3*n+1,wg,length)
        }
        wg.Done()
    }()
}

func main()  {

    var wg sync.WaitGroup
    ans := answer{0,0}
    length := 0

    for i := 1; i < max; i++ {
        length = 0
        collatz(i,&wg,&length)
        wg.Wait()
        if length > ans.length {
            ans.value = i
            ans.length = length
        }
    }

    fmt.Println(ans)
}
