package main

import (
	"fmt"
	"time"
	"math"
	"sort"
)
type int64arr []int64
func (a int64arr) Len() int { return len(a) }
func (a int64arr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a int64arr) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	t0 := time.Now()
	var a int64arr
	//32 and 20 are magic numbers
	//this loop misses two numbers smaller than the biggest one it finds (bigger i, smaller j)
	//so I had to compute a(32) to properly find a(30)
	for i := 7; len(a) < 32 ; i++ {
		for j := 2;j <= 20; j++ {
			tmp := int64(math.Pow(float64(i),float64(j)))
			if int64(i) == digitSum(tmp) {
				//fmt.Println(tmp)
				a = append(a,tmp)
				//fmt.Println(len(a),i,j)
			}
		}
	}
	sort.Sort(a)
	fmt.Println("Answer: ",a[20], "Time: ",time.Since(t0))
}

func digitSum(n int64) (s int64) {
	for n > 0 {
		s += (n % 10)
		n /= 10
	}
	return
}