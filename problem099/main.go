package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	t0 := time.Now()
	ans := findAnswer()
	fmt.Println("Answer:", ans, "\nTime:", time.Since(t0))

}

func findAnswer() int {
	values := parseInput("p099_base_exp.txt")
	max := float64(0)
	lineNum := 1
	for i := 1; i < len(values)-1; i++ {
		if values[i-1] > max {
			lineNum = i
			max = values[i-1]
		}
	}
	return lineNum
}

func parseInput(path string) (values []float64) {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		a, _ := strconv.Atoi(line[0])
		b, _ := strconv.Atoi(line[1])
		values = append(values, evaluate(a, b))
	}
	return
}

//Comparing the logarithms of the numbers maintains the ordering while preventing any sort of overflow
func evaluate(n, m int) float64 {
	return float64(m) * math.Log(float64(n))
}
