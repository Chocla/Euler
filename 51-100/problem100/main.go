package main

import (
	"fmt"
	"time"
)

const (
	max = 1000000000000
)

func main() {
	t0 := time.Now()
	ans := findAnswer()
	fmt.Println("Answer:", ans, "\nTime:", time.Since(t0))

}

func findAnswer() int64 {
	x := int64(1) //blue
	y := int64(1) //total
	for y < 2*max-1 {
		x, y = 2*y+3*x, 3*y+4*x
	}
	return (x - 1) / 2
}
