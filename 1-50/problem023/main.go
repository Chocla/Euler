package main

import (
	"fmt"
)

func main(){
	answer := 0
	//get a list of abundant numbers less than 20161
	abNums := getAbNums()
	
	for i := 1; i < 28123; i++ {
		flag := true
		for _,v := range abNums{
			if v > i {
				break
			}
			if isIn(i - v,abNums) {
				flag = false
				break
			}
		}
		if flag {
			answer += i
		}
	}
	fmt.Println(answer)
}
func isIn(n int, s []int)(bool){
	for _,v := range s {
		if v == n {
			return true
		}
	}
	return false
}
func getAbNums()(n []int) {
	for i := 0; i < 28123; i++ {
		divSum := 0
		for j := 1; j < i; j++ {
			if i % j == 0 {
				divSum += j
			}
		}
		if divSum > i {
			n = append(n, i)
		}
	}
	return
}