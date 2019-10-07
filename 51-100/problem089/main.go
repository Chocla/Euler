package main

import (
	"time"
	"bufio"
	"os"
	"fmt"
)
var numericalValues = []int{
	1000,900,500,400,100,90,50,40,10,9,5,4,1,
}
var romanValues = []string{
	"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I",
}

func main() {
	t0 := time.Now()
	ans := findAnswer("p089_roman.txt")
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
}
func findAnswer(path string) (ans int){
	oStrings := parseInput(path)
	ans = calculateLengthDiff(oStrings)
	return
}

func calculateLengthDiff(oStrings []string) (diff int) {
	for i := range(oStrings) {
		len1 := len(oStrings[i])
		len2 := len(intToMinimalNumeral(numeralToInt(oStrings[i])))
		diff += (len1 - len2)
	}
	return
}

func parseInput(path string) (originalStrings []string) {
	f,err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		originalStrings = append(originalStrings, scanner.Text())
	}
	return
}

func numeralToInt(s string) (n int){
	var i int //keep in scope after loop finishes
	for i = 0; i < len(s) - 1; i++ {
		substr1 := s[i:i + 2]
		substr2 := s[i:i + 1]
		val,f := checkSubStr(substr1,substr2)
		if f {i++} //double character found
		n += val

	}
	substr3 := s[i:] 
	val,_ := checkSubStr(substr3,"")
	n+= val

	return
}
func checkSubStr(s1,s2 string) (int,bool){
	for j := range(romanValues) {
		if s1 == romanValues[j] {
			return numericalValues[j],true
		}
	}
	for j := range(romanValues) {
		if s2 == romanValues[j] {
			return numericalValues[j],false
		}
	}
	return 0,false
}
func intToMinimalNumeral(n int) (s string){
	for i := range numericalValues {
		for n - numericalValues[i] >= 0 {
			s += romanValues[i]
			n -= numericalValues[i]
		}
	}
	return
}