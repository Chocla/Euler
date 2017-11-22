package main

import (
	"fmt"
	"math"
	"strconv"
)
//Find the smallest cube for which exactly five permutations of its digits are cube.


func main() {
	//fmt.Println(findAllPermutations([]int{1,0,3,0,3,0,1}))
	
	 i := 0
	answerFound := false;
	for ;!answerFound; i++ {
		answerFound = checkPermutationsforCubes(int(math.Pow(float64(i),3)))
		fmt.Println(i)
	}
	fmt.Println("i:", i)
	fmt.Println("Answer:", i)
}

func checkPermutationsforCubes(inCube int)(x bool){
	//fmt.Println(inCube)
	digits := getDigits(inCube)
	//fmt.Println(digits)
	numCubes := findAllPermutations(digits)
	//testing case
	//fmt.Println(numCubes)
	if numCubes == 5 {
		return true
	}
	return false
}

//Implementation of heap's algorithm for permutations
func findAllPermutations(input []int) (numCubes int){
	c := make([]int, len(input))
	cubes := make([]int, 0)

	for i := 0; i < len(input); {
		if c[i] < i {
			if i % 2 == 0 {
				swap(&input[0], &input[i])
			} else {
				swap(&input[c[i]], &input[i])
			}
			num,isCube := checkCube(input)
			if isCube{
				uniqueCube := true
				for i := 0; i < len(cubes);i++{
					if num == cubes[i] {
						uniqueCube = false
					}
				}
				if uniqueCube {
					cubes = append(cubes, num)
				}
			}
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i += 1
		}
	}
	//fmt.Println(cubes)
	return len(cubes)
}
func swap(a, b *int) {
	var tmp int = *a
	*a = *b
	*b = tmp

}
func getDigits(input int)(digitSlice []int) {
	x := strconv.Itoa(input)
	//fmt.Println(x)
	for i := 0; i < len(x); i++ {
		digitSlice = append(digitSlice, int(x[i] - '0'))
	}

	return
}

func checkCube(input []int) (num int, x bool) {
	//convert back into one number

	place := 1
	for i := len(input) - 1; i >= 0; i-- {
		num += input[i]*place
		place *= 10
	}
	if num < place / 10 {
		return 0,false
	}
	cbrt := math.Cbrt(float64(num))
	//fmt.Println(cbrt,math.Floor(cbrt))
	if cbrt - math.Floor(cbrt) < 0.0000000000000000000000001 {
		return num,true
	}
	return 0,false
	//check the cuberoot
		//if it's a perfect cube, return true
}