package main

import (
	"fmt"
	"time"
)
const (
	jan = 31
	mar = 31
	apr = 30
	may = 31
	jun = 30
	jul = 31
	aug = 31
	sep = 30
	oct = 31
	nov = 30
	dec = 31

)
func main(){
	sunCount := 0
	t0 := time.Now()
	for year := 1900 ; year < 2000; year++ {
		for month := time.January; month <= 12; month++ {
			if time.Date(year,month,1,1,1,1,1,time.UTC).Weekday() == time.Sunday {
				sunCount++
			}
		}
	}
	t := time.Since(t0)	
	fmt.Println("Answer: ", sunCount, "\nin ",t)
}