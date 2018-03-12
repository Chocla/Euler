package main

import (
	"fmt"
	"math/big"
	"strings"
	"strconv"
)

func main() {
	answer := 0
	exponent := big.NewInt(0)
	x := big.NewInt(2)
	y := big.NewInt(1000)
	m := big.NewInt(0)
	exponent.Exp(x,y,m)
	str := exponent.String()
	strarray := strings.Split(str, "")
	
	
	for i := 0; i < len(strarray); i++ {
		tmp,_ := strconv.Atoi(strarray[i])
		answer += tmp
	}
	fmt.Println(answer)
}
