package main

import (
	"fmt"
)
func main(){
	answer := 1;
	for i := 1001; i >= 3; i -= 2 {
		answer += (4*i*i - 12*i + 14)
	}
	fmt.Println("Answer: ", answer)
}
