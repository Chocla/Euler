package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
	"sort"
)

func main() {
	answer := 0
	var names []string
	file,err := os.Open("p022_names.txt")
	if(err != nil) {
		panic(err)
	}
	reader := bufio.NewReader(file)
	_, err = reader.Discard(1)

	//put names into string array
	for {
		tmpName,err := reader.ReadString('"')
		tmpName = strings.TrimRight(tmpName, "\"")
		names = append(names, tmpName)
		_,err = reader.Discard(2)
		if err == io.EOF {
			break
		}
	}
	//sort names into alphabetical order
	sort.Strings(names)

	//determine name score
	for i := 0; i < len(names);i++ {
		tmp := computeNameScore(names[i],i+1)
		fmt.Println("Name: ",names[i], "Score: ",tmp)
		answer += tmp
	}
	fmt.Println("Total: ",answer)
}
func computeNameScore(name string, numInList int)(n int) {
	for i := 0; i < len(name); i++ {
		n += int(name[i]) - 64
	}
	return numInList*n
}