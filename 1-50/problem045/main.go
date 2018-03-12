package main

import (
	"math"
	"fmt"
)
//since the triangular numbers are a subset of the hexagonal, we only need to find a num that is both pentagonal and hexagonal
//generate pentagonal numbers, and check if they are hexagonal
func main() {
	ans := problem45()
	fmt.Println(ans*(3*ans-1)/2)
}

func problem45() ( int) {
	for i := 166; ; i++ {
		if isHex(genPent(i)) {
			return i
		}
	}

}

func genPent(n int) (pn int64 ){
	a := int64(n)
	return a*(3*a-1)/2
}

//if (sqrt(8x+1)+1)/4 is a natural number, x is hexagonal
func isHex(x int64) (bool){
	n := (math.Sqrt(8*float64(x)+1) + 1) / 4
	if n == float64(int64(n)) {
		return true
	}
	return false
}