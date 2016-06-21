package main

import (
	"fmt"
	"math/big"
)

func main(){
	answer := big.NewInt(0)
	for i := 1; i <= 1000; i++ {
		bigi := big.NewInt(int64(i))
		tmp := big.NewInt(0)
		answer.Add(answer,tmp.Exp(bigi,bigi,tmp))
	}
	fmt.Println(answer)
}