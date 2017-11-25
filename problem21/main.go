package main

import ( 
	"fmt"
)
type amicable struct {
	n int
	dn int
	flag bool
}
func main(){
	answer := 0
	x := make([]amicable,0)
	//fill array 
	for i := 2; i < 10000; i++ {
		x = append(x, amicable{i, dn(i),false })
	}
	// go through array and check for amicability
	for i := 2; i < len(x); i++ {
		
		if x[i].flag || x[i].dn == 1 || x[i].dn-2 > len(x){
			continue
		}
		if isAmicable(x[i], x[x[i].dn-2]) {
			answer += x[i].n + x[i].dn
			x[x[i].dn-2].flag = true
			fmt.Println(x[i],x[x[i].dn-2])
			
		}
		
	}	
	fmt.Println("answer:", answer)
}

func isAmicable(a, b amicable) (bool) {
	if a.n == b.dn && b.n == a.dn && a.n != b.n{
		return true
	}
	return false
}
func dn(n int)(x int) {
	for i := 1; i < n; i++ {
		if n % i == 0 {		
			x += i
		}
	}
	
	return
}