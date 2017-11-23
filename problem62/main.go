package main

import (
	"fmt"
	"strconv"
	"sort"
)
//Find the smallest cube for which exactly five permutations of its digits are cube.

type cube struct {
	o []int // ordered digits
	t int   // times appeared
	l int   // lowest root
}
func main() {
	answer := 0
	i := 0
	var cubeList []cube
	initCube := cube{[]int{0},1,0}
	cubeList = append(cubeList, initCube)
	answerFound := false

	for !answerFound {
		i++
		tmpCube := orderDigits(i*i*i)
		inList := false

		for j := 0; j < len(cubeList); j++ {

			for k := 0; k < len(tmpCube); k++ {
				if testEq(cubeList[j].o, tmpCube) {
					inList = true
				}
			}
			if inList {
				cubeList[j].t++
				break
			}


		}
		if !inList {
			newCube := cube{tmpCube,1,i}
			cubeList = append(cubeList, newCube)
		}

		for j := 0; j < len(cubeList);j++ {
			if cubeList[j].t ==5 {
				fmt.Println(cubeList[j])
				answer = cubeList[j].l
				answerFound = true
			}
		}

	}
	fmt.Println("Answer:", answer*answer*answer)
}

func testEq(a, b []int) bool {
	
		if a == nil && b == nil { 
			return true; 
		}
	
		if a == nil || b == nil { 
			return false; 
		}
	
		if len(a) != len(b) {
			return false
		}
	
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
	
		return true
	}
func orderDigits (i int) (o []int) {
	var digitSlice sort.IntSlice
	x := strconv.Itoa(i)

	for j := 0; j < len(x); j++ {
		digitSlice = append(digitSlice, int(x[j] - '0'))
	}

	digitSlice.Sort()
	o = digitSlice

	return
}


