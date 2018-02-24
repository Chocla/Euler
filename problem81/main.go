package main

import (
	"time"
	"strconv"
	"strings"
	"bufio"
	"os"
	"fmt"
)
var test = [][]int{  {131,673,234,103,18},
					 {201,96,342,965,150},
					 {630,803,746,422,111},
					 {537,699,497,121,956},
					 {805,732,524,37,331}}
func main() {
	t0 := time.Now()
	ans := findAnswer()
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
}

func findAnswer() (ans int) {
	mat := parseInput("p081_matrix.txt")
	return findMinPath(mat)
}
func parseInput(path string) (matrix [][]int){
	file,err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		tmp := make([]int,0)
		line := strings.Split(scanner.Text(),",")
		for i := range line {
			n,_ := strconv.Atoi(line[i])
			tmp = append(tmp,n)
		}
		matrix = append(matrix,tmp)
	}


	return
}

func buildPSumArray(n,m int) ([][]int) {
	a := make([][]int,n)
	for i := range a {
		a[i] = make([]int,m)
	}
	return a
}
func findMinPath(mat [][]int) (int) {
	pSums := buildPSumArray(len(mat),len(mat[0]))
	pSums[0][0] = mat[0][0]
	//edges only have one option, fill those in first
	for i := 1; i < len(mat);i++ {
		pSums[0][i] += pSums[0][i-1] + mat[0][i]
		pSums[i][0] += pSums[i-1][0] + mat[i][0]
	}
	for i := 1; i < len(mat); i++ {
		for j := 1; j < len(mat); j++ {
			pSums[i][j] = mat[i][j] + min(pSums[i-1][j],pSums[i][j-1])
		}
	}
	return pSums[len(mat)-1][len(mat[0])-1]
}

func min(a,b int) (int){
	if a < b {
		return a
	}
	return b
}
