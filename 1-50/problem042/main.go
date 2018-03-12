package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
	"time"
)
const inputPath = "p042_words.txt"

func main() {
	t0 := time.Now()
	ans := findAnswer(inputPath)
	fmt.Println("Answer: ",ans, "Time: ", time.Since(t0))
}
func findAnswer(path string) (ans int) {
	words := parseInput(path)
	triNums := findTriangleNumbers(300) //300 arbitrary
	for _,word := range words {
		score := findWordScore(word)
		for _, n := range triNums {
			if score == n {
				ans++
				break
			}
		}
	}
	return
}
func findTriangleNumbers(max int) (triNums []int) {
	for i := 0; i <= max; i++ {
		triNums = append(triNums, (i * (i+1))/2)
	}
	return
}
func findWordScore(word string) (score int) {
	for _,b := range word {
		//ASCII codes for capital letters start at 65
		//subtracting 64 calculates their position in the alphabet
		score += (int(b) - 64)
	}
	return
}
func parseInput(path string) (words []string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = strings.Split(scanner.Text(), "\",\"")
		words[0] = strings.TrimLeft(words[0], "\"")
		words[len(words)-1] = strings.TrimRight(words[len(words)-1],"\"")
		//fmt.Println(words)
	}
	return words
}