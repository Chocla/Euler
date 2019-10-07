package main

import (
	"fmt"
	"math"
	"strings"
	"os"
	"bufio"
)

func main() {
	uBound := 1000000
	primes := seive(uBound)
	sVals := populateValues(uBound,primes)
	s := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sVals)), ","), "[]")
	f, err := os.Create("goSigDump.txt")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	defer f.Close()
	w.WriteString(s)
	fmt.Println()
}
func seive(upperBound int) (primes []int) {
	nums := make([]bool,upperBound)
	for i := 2; i < upperBound;i++ {
		if !nums[i] {
			primes = append(primes,i)
			//nums[i] = true
			for j := 2*i; j < upperBound; j+= i {
				nums[j] = true
			}
		}
	}

	return
}
func sigma(n int,primes []int) (sig int) {
	sig = 1
	tmpN := n
	for i := range primes {
		if n == primes[i] {
			return 1
		}
		if n < primes[i] {
			break
		}
	}
	di := 0
	for tmpN > 1 {

		p := primes[di]
		a := 0
		for tmpN % p == 0 {
			tmpN = tmpN / p
			a++
		}
		sig *= ((int(math.Pow(float64(p),float64(a+1))) - 1) / (p - 1))
		di++
	}

	return sig - n
}
func populateValues(size int,primes []int) ([]int) {
	sVals := make([]int, size+1)
	for i := range sVals {
		if i % 10000 == 0 {
			fmt.Println(i)
		}
		sVals[i] = sigma(i,primes)
	}
	return sVals
}
