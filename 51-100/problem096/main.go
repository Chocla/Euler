package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main(){
	
	file,err := os.Open("p096_sudoku.txt")
	if(err != nil) {
		panic(err)
	}
	reader := bufio.NewReader(file)
	puzzleCount := 0
	answer := 0
	for i := 0; i < 50; i++{
		puzzleCount++
		if i == 0 {
			_, err = reader.Discard(8)
		} else {
			_,err = reader.Discard(7)
		}
		if(err != nil){
			panic(err)
		}
		str,_ := reader.ReadString('G')
		
		puzzle := fillPuzzle(strings.TrimRight(str,"\nG"))
		answer += solvePuzzle(puzzle)
		
	}
	fmt.Println("Answer: ", answer)

}

func fillPuzzle(input string) (puzzle [9][9]int) {
	rows := strings.SplitAfter(input, "\n")
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
	 		puzzle[i][j] = int(rows[i][j] - '0') //magic byte to int conv
		}
	}
	return
}
func isValid(candidate, row, col int , p[9][9]int)(v bool) {
	v = true
	//check row
	for i := 0; i < 9; i++ {
		if i == col {
			continue
		}
		if p[row][i] == candidate {
			v = false
			break
		}
	}
	//check col
	for i := 0; i < 9; i++ {
		if i == row {
			continue
		}
		if p[i][col] == candidate {
			v = false
			break
		}
	}
	//check box
	boxRow := 3 * (row / 3)
	boxCol := 3 * (col / 3)
	for i := boxRow; i < boxRow + 3 ; i++ {
		for j := boxCol; j < boxCol + 3 ; j++ {
			if (i != row || j != col )&& p[i][j] == candidate {
					v = false
					break

			}

		}
	}
	return v
}
func solvePuzzle(p [9][9]int) (partialSum int){
	test := new([81]int)
	for i := 0; i < 9; i++{
		for j := 0; j < 9; j++ {
			test[9*i + j] = p[i][j]
		}
	}
	stepsBack := new([81]int)
	tmp := 0
	for i := 0; i < 81; i++{
		
		stepsBack[i] = 1
		if test[i] != 0 {
			tmp++
			continue
		}
		stepsBack[i] += tmp
		tmp = 0
	}

	currentCell := 0
	//brute force algorithm 
	for currentCell < 81 {
		if test[currentCell] > 0 {
			currentCell++ 
		}  else {
			row := currentCell / 9
			col := currentCell % 9
			p[row][col] += 1
			if p[row][col] > 9 {
				p[row][col] = 0
				currentCell -= stepsBack[currentCell]
				continue
			}
			if (isValid(p[row][col], row, col, p)){
				currentCell++;
			}
		}
	}
	partialSum = 100*p[0][0] + 10*p[0][1] + p[0][2]
	return;
}