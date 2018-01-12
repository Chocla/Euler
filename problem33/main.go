package main

import (
	"fmt"
)

func main() {

	problem33()	
}

func problem33() (ans int) {
	for num := 11; num < 100; num++ {
		if num % 10 == 0 {
			continue
		}
		for den := num+1; den < 100; den++ {
			if den % 10 == 0 {
				continue
			}
			numerDigits := intToDigits(num)
			denomDigits := intToDigits(den)
			numerDigits,denomDigits,f := cancelDigits(numerDigits,denomDigits)
			
			a := float64(num) / float64(den)
			b := float64(digitsToInt(numerDigits)) / float64(digitsToInt(denomDigits))
			
			if a == b && f{
				//manually find answer from this point (multiply and simplify 4 fractions)
				fmt.Printf("%v/%v ",num,den)
			}
		}
	}

	return
}

func cancelDigits(a,b []int) (c,d []int, f bool) {
	for k,v := range a {
		for k2,v2 := range b {
			if v == v2 {
				f = true
				a[k],b[k2] = -1,-1
				break
			}
		}
	}
	for i := range a {
		if a[i] != -1 {
			
			c = append(c,a[i])
		}
	}
	for i := range b {
		if b[i] != -1 {
			d = append(d,b[i])
		}
	}
	return
}
func digitsToInt(d []int) (n int) {
	for i,mul := len(d) - 1,1; i >= 0; i,mul = i-1, mul * 10 {
		n += d[i]*mul
	}
	return
}
func intToDigits(n int) (d []int) {
	for ; n > 0; n /= 10 {
		d = append([]int{n % 10}, d...)
	}
	return
}

