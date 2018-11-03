package main

import (
	"strings"
	"fmt"
	"math/big"
	"time"
)

func main(){
	t0 := time.Now()
	ans := findAnswer()
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
}

func findAnswer()(ans int) {
	for i := 1; i <100; i++ {
		ans += digitSum(generateRoot(i))
	}
	return 
}
func generateRoot(n int) string{
	root := big.NewFloat(float64(n))
	root.SetPrec(1000)
	root.Sqrt(root)

	//Don't count perfect squares
	if root.IsInt() {
		return "0"
	}
	//Convert to string containing first 100 digits
	rootS := root.Text('f',1000)
	rootS = strings.Join(strings.Split(rootS,"."),"") 
	rootS = rootS[0:100]
	return rootS

}

func digitSum(s string) (sum int){
	for _,v := range(s) {
		sum += int(v - '0')
	}
	return
}