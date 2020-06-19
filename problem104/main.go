package main

import (
	"fmt"
	"math/big"
	"sort"
	"time"
)

type sortRunes []rune 
func (s sortRunes) Less(i, j int) bool { return s[i] < s[j] }
func (s sortRunes) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s sortRunes) Len() int { return len(s) }

func SortString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}

func main() {
	t0 := time.Now()
	ans := findPandigitalFib()
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
	
}

func findPandigitalFib() (int){

	a := big.NewInt(1)
	b := big.NewInt(1)

	for i := 3; ; i++ {
		a.Add(a,b)
		a, b = b, a
		str := b.String()
		if i > 45{
			last := SortString(str[len(str) - 9:])
			first := SortString(str[:9])
			if last == "123456789" && first == "123456789" {
				return i
			}
		}	
	}
	return -1
}