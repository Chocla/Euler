package main

import (
	"fmt"
)

func main() {
	count := 0
	//WHO NEEDS CLEVER ANSWERS
	for a := 0; a <= 1; a ++ {
		for b := 0; b <= 2; b++ {
			for c := 0; c <= 4; c++ {
				for d := 0; d <= 10; d++ {
					for e := 0; e <= 20; e++ {
						for f := 0; f <= 40; f++ {
							for g := 0; g <= 100; g++ {
								for h := 0; h <= 200; h++ {
									if (a*200 + b*100 + c*50 + d*20 + e*10 + f*5 + g*2 + h) == 200 {
										count++
									}
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(count)
}