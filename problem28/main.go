package main

import (
	"fmt"
)


func main(){
	answer := 0;
	for i := 1001; i >= 3; i -= 2 {
		answer += i*i + (i-1)*(i-1) + (i-2)*(i-2) + (i-3)*(i-3)
	}
	answer += 1
	fmt.Println("Answer: ", answer)
}
