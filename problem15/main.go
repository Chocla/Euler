package main

import (
	"fmt"
)
//the total number of unique NE paths from (0,0) to (a,b) is 
//(a+b)!/(a)!(a!)
func main() {
	//took the time to simplify 40 choose 20
	fmt.Println(4*39*37*35*33*31*29*23)
}

