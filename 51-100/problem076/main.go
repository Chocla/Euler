package main

import(
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	ans := findAnswer()
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
}

func findAnswer() (int) {
	part := []int{1}
	pents := []int{0}
	j := 1
	for i := 1; i < 101; i++ {
		if i + 1> pents[len(pents) - 1] {
			pents = append(pents,nextPent(j)...)
			j++
		}

		part = append(part, nextPart(part,pents))
	}

	//p(n) counts the sum "n" as a valid partition
	return part[100] - 1
}

func nextPart(prevParts []int, pents []int) (pn int) {
	sign := 1
	n := len(prevParts)
	for i := 1; (n >= pents[i]) ; i++ {
		term := prevParts[n - pents[i] ]
		pn += sign*term

		if i % 2 == 0 && i != 0{
			sign *= -1
		}
	}
	return pn
}
//calculates the next two generalized pentagonal numbers (using n and -n)
func nextPent(n int) ([]int) {
	return []int{(3*n*n -n)/2, (3*n*n + n)/2 }
}