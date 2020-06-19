package main

import (
	"fmt"
	"time"
	"math"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	t0 := time.Now()
	ans := findAnswer()
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
}

func findAnswer() (s int){
	sums := make([]int,10)
	maxCount := make([]int,10)

	for i := 6; i < 47; i++ {
		path := "primes\\2T_part" + strconv.Itoa(i) + ".txt"
		primes := parsePrimes(path)
		for _,p := range primes {
			dc := digitCount(p)
			for i,count := range dc {
				// new maximum, reset sum to be p
				if count == maxCount[i] {
					sums[i] += p
				}
				if count > maxCount[i] {
					maxCount[i] = count
					sums[i] = p
				}
			}
		}
	}
	for _,v := range(sums) {
		s += v
	}
	return  
}
func parsePrimes(path string) (primes []int) {
	file,err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := strings.Split(scanner.Text(),"\t")
		for _,v := range(str) {
			if len(v) == 10 {
				tmp,_ := strconv.Atoi(v)
				primes = append(primes,tmp)
			}
		}
	}
	return
}
func digitCount(p int) (map[int]int) {
	dc := make(map[int]int)
	for p > 0 {
		dc[p%10]++
		p /= 10
	}

	return dc
}
