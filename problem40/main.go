package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//8.5 Minutes.. Yikes
	t0 := time.Now()
	champ := ""
	for i,length := 1,0; length < 1000000;i++ {
		t := strconv.Itoa(i)
		length += len(t)
		champ += t
	}

	answer := (champ[0]-'0') * (champ[9] - '0') * (champ[99] - '0') * (champ[999] - '0') * (champ[9999] - '0') * (champ[99999] - '0') * (champ[999999] - '0')
	fmt.Println("Answer: ",answer, "Time: ",time.Since(t0))
}