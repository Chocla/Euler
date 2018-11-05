package main

import (
	"fmt"
	"time"
)

//we dont need to generate squares of primes past the square root of the biggest value
const sqrt_max = 11243247

func main() {
	t0 := time.Now()
	ans := findAnswer(51)
	fmt.Println("Answer:", ans, "\nTime:", time.Since(t0))
}

func findAnswer(depth int) int {
	candidates := removeDuplicates(pascal(depth))
	squares := primeSquareSeive(sqrt_max)
	return countSquareFreeValues(candidates, squares)
}
func countSquareFreeValues(candidates, squares []int) (sum int) {
	for i := range candidates {
		flag := true
		for j := range squares {
			if candidates[i]%squares[j] == 0 {
				flag = false
				break
			}
		}
		if flag {
			sum += candidates[i]
		}
	}
	return
}

//generates pascal's triangle, then concatenates the first half of each row into one slice
func pascal(depth int) (triConcat []int) {
	triConcat = append(triConcat, []int{1, 2}...)
	triangle := make([][]int, 0)
	triangle = append(triangle, []int{1})
	for i := 1; i < depth; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
		triangle = append(triangle, row)
	}
	for i := 3; i < len(triangle); i++ {
		triConcat = append(triConcat, triangle[i][1:i/2+1]...)
	}
	return
}

//Sieve of Eratosthenes that returns squares of the primes instead
func primeSquareSeive(max int) (p []int) {
	flags := make([]bool, max+1)
	flags[0], flags[1] = true, true
	for i := 2; i < len(flags); i++ {
		if !flags[i] {
			for j := 2 * i; j < len(flags); j += i {
				flags[j] = true
			}
		}
	}

	for i := range flags {
		if !flags[i] {
			p = append(p, i*i)
		}
	}

	return
}

func removeDuplicates(elements []int) []int {
	encountered := map[int]bool{}
	result := []int{}
	for v := range elements {
		if encountered[elements[v]] {
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}
