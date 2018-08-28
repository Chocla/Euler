package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	ans := findAnswer(50)
	//fmt.Println(findMax(7,5,6))
	fmt.Println("Answer:",ans,"\nTime:",time.Since(t0))
}

func findAnswer(n int) (ans int){
	for x1 := 0; x1 <= n; x1++ {
		for y1 := 0; y1 <= n; y1++ {
			for x2 := 0; x2 <= n; x2++ {
				for y2 := 0; y2 <= n; y2++ {
					if x1 == x2 && y1 == y2 {
						continue; //same point
					} 
					if (x1 == 0 && y1 == 0 ) || (x2 == 0 && y2 == 0){
						continue;
					}
					a,b,c := genSides(x1,x2,y1,y2)

					switch findMax(a,b,c) {
					case 1: //a max
						if a == b + c {
							ans++
						}
						break
					case 2: //b max
						if b == a + c {
							ans++
						}
						break
					default:
						if c == a + b {
							ans++

						}
					}
				}
			}
		}
	}
	return ans/2
}
func genSides(x1,x2,y1,y2 int) (a,b,c int) {
	return x1*x1 + y1*y1, x2*x2+y2*y2, (x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)
}
func findMax(a,b,c int) (int) {
	if a > c {
		if b > a {
			return 2
		} 
		return 1
	 }	else if b > c {
		return 2
	}
	return 0
}