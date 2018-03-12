/*Problem 13*/
package main

import (
	"fmt"
	"os"
	"bufio"
	"math/big"
)
const (
	length int = 100
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
func main()  {
	answer := big.NewInt(0)
	
	file, err := os.Open("in.txt")
	checkErr(err)
	reader := bufio.NewReader(file)
    for i := 0; i < length; i++ {
	    tmpBigInt := big.NewInt(0)
		tmp,_ := reader.ReadString('\n')
		tmpBigInt.SetString(tmp,10)
		answer.Add(answer,tmpBigInt)
	}

	fmt.Printf("Answer = %v\n",answer)
}
