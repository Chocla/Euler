package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	t0 := time.Now()
	ans := parseInput("p102_triangles.txt")
	fmt.Println("Answer: ", ans, "\nTime:", time.Since(t0))
}
func parseInput(path string) (ans int) {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		coords := make([]int, 0)
		for _, v := range line {
			n, _ := strconv.Atoi(v)
			coords = append(coords, n)
		}
		if originInside(coords) {
			ans++
		}
	}
	return
}
func originInside(coords []int) bool {
	return area(coords[0], coords[1], coords[2], coords[3], coords[4], coords[5]) == areaTwoPoints(coords[0], coords[1], coords[2], coords[3])+
		areaTwoPoints(coords[0], coords[1], coords[4], coords[5])+
		areaTwoPoints(coords[2], coords[3], coords[4], coords[5])

}
func area(x1, y1, x2, y2, x3, y3 int) float64 {
	return .5 * abs((x1-x3)*(y2-y1)-(x1-x2)*(y3-y1))
}
func areaTwoPoints(x1, y1, x2, y2 int) float64 {
	return .5 * abs(x1*y2-x2*y1)
}

func abs(x int) float64 {
	if x > 0 {
		return float64(x)
	}
	return -float64(x)
}
