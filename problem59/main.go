package main

import (
	"time"
	"strconv"
	"strings"
	"bufio"
	"os"
	"fmt"
)
const path = "p059_cipher.txt"

func main() {
	fmt.Println()
	t0 := time.Now()
	msg := parseInput(path)
	ans := genKeys(msg,97,122)
	fmt.Println("Answer: ", ans, "\nTime:",time.Since(t0))
}
func countVals(msg []byte) (count int) {
	for _,v := range msg {
		count += int(v)
	}
	return
}
func genKeys(code []byte, min,max int) (int) {
	for i := min; i <= max; i++ {
		for j := min; j <= max; j++ {
			if j == i {
				continue
			}
			for k := min; k <= max; k++ {
				if k == i || k == j {
					continue
				}
				candidate := decode(code,[]byte{byte(i), byte(j), byte(k)})
				if analyzeString(candidate) {
					fmt.Println(candidate)
					return countVals([]byte(candidate))
					
				}
			}
		}
	}
	return 0
}
func analyzeString(s string)(bool){
	the := " the "
	for i := range s {
		if i > len(s) - 5 {
			break
		}
		if strings.EqualFold(s[i:i+5],the) {
			return true
		}
	}
	return false
	//if string contains the substring " the "
		//print out full string for closer inspection
}

func decode(code []byte, key []byte )(string) {
	newMsg := make([]byte,len(code))
	for i := range code {
		newMsg[i] = code[i] ^ key[i % 3]
	}

	return string(newMsg)
}
func parseInput(path string) ([]byte){
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	scan.Scan()
	line := strings.Split(scan.Text(),",")
	codes := make([]byte, len(line))
	for i := range line {
		tmp, err := strconv.Atoi(line[i])
		if err != nil {
			panic(err)
		}
		codes[i] = byte(tmp)
	}
	
	return codes
}