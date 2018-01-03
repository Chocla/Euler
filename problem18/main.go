package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"time"
)
const inputPath = "problem67.txt"

func main() {

	t0 := time.Now()
	pyramid := parseInput(inputPath)
	ps      := makePartialSumArray(len(pyramid))
	ps[0][0] = pyramid[0][0]
	//Partial Sums at each step in the pyramid
	for i := 1; i < len(pyramid);i++ {
		for j := 0; j < len(pyramid[i]); j++ {
			//partial sum contains pyramid value
			ps[i][j] += pyramid[i][j]
			//if position is on the edge, only one value can be added
			//else, add the max of the two adjacent values
			switch j {
			case 0: ps[i][j] += ps[i-1][j]
			case len(pyramid[i]) - 1 : ps[i][j] += ps[i-1][j-1] 
			default: ps[i][j] += max(ps[i-1][j], ps[i-1][j-1])
			}
		}
	}
	//maximum value in final row is the answer
	ans := 0
	lRow := len(ps) - 1
	for i := range ps[lRow] {
		if ps[lRow][i] > ans {
			ans = ps[lRow][i]
		}
	}
	fmt.Println("Answer: ",ans, "Time: ",time.Since(t0))
}
func max(a,b int) (int) {
	if a > b {
		return a
	}
	return b
}
func makePartialSumArray(len int) (ps [][]int) {
	for i := 0; i < len; i++ {
		ps = append(ps, make([]int,i+1))
	}
	return
}
func parseInput(path string) (p [][]int) {
	file,err := os.Open(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var tmp []int
		line := strings.Split(scanner.Text()," ")
		for i := range line {
			n,_ := strconv.Atoi(line[i])
			tmp = append(tmp,n)
		}
		p = append(p,tmp)
	}
	return
}