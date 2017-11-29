package main

import (
	"fmt"
	"time"
)
const size  = 1001
const total = size*size

func main(){
	if size % 2 != 1 {
		fmt.Println("size must be odd")
		return
	}

	var spiral [size][size]int
	t0 := time.Now()
	fillSpiral(&spiral)
	answer := getSum(&spiral)
	t1 := time.Now()
	
	fmt.Println("Answer: ", answer , " in ", t1.Sub(t0))
}
func fillSpiral(a *[size][size]int) {
	//start in the middle
	row := size / 2
	col := size / 2
	num := 1
	dir := 0
	stepSize := 1
	a[row][col] = num
	for num < total - size{
		for i := stepSize; i > 0; i-- {
			num++
			switch dir{
			case 0: col++ 
				break
			case 1: row++ 
				break
			case 2: col-- 
				break
			case 3: row--
				break
			}
			a[row][col] = num
		}
		dir = (dir + 1) % 4
		for i := stepSize; i > 0; i-- {
			num++
			switch dir{
			case 0: col++ 
				break
			case 1: row++ 
				break
			case 2: col-- 
				break
			case 3: row--
				break
			}	
			a[row][col] = num
		}
		dir = (dir + 1) % 4
		stepSize++
	}
	for i := 1; i < size; i++ {
		num++
		a[0][i] = num
	}

}

func getSum(a *[size][size]int) (sum int) {
	for i,j := 0,size-1; i < size && j >= 0; i,j = i+1,j-1 {
		if i != j {
			//fmt.Println(a[i][i] , a[i][j])
			sum += (a[i][i] + a[i][j])
		} else {
			//fmt.Println(a[i][i])
			sum += a[i][i]
		}
	}
	return
}