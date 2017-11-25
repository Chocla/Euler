/*Problem 17
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?
*/
package main

import (
	"fmt"
)
/*
one         3
two         3
three       5
four        4
five        4
six         3
seven       5
eight       5
nine        4
ten         3
eleven      6
twelve      6
thirteen    8
fourteen    8
fifteen     7
sixteen     7
seventeen   9
eighteen    8
nineteen    8
twenty      6
thirty      6
forty       5
fifty       5
sixty       5
seventy     7
eighty      6
ninety      6
*hundred    7
thousand    8
and         3
*/

func main()  {
	answer := 0
	//answer += 119 % 100
	for i := 1; i <= 1000; i++ {
		tmp := intToLetter(i)
		fmt.Println("i: ", i, "letters:",tmp)
		answer += tmp
	}
	//answer += intToLetter(342)
	fmt.Println("Answer: ", answer)
}

func intToLetter(x int) (letterCount int) {
	onesPlace := (x - (x /10 )*10)
	tensPlace := (x - onesPlace - (x /100)*100) /10
	hundredsPlace := (x - onesPlace - tensPlace*10) /100
	switch x % 100 {
		case 11:   return countOnes(hundredsPlace) + 10 + 6
		case 12:   return countOnes(hundredsPlace) + 10 + 6
		case 13:   return countOnes(hundredsPlace) + 10 + 8
		case 14:   return countOnes(hundredsPlace) + 10 + 8
		case 15:   return countOnes(hundredsPlace) + 10 + 7
		case 16:   return countOnes(hundredsPlace) + 10 + 7
		case 17:   return countOnes(hundredsPlace) + 10 + 9
		case 18:   return countOnes(hundredsPlace) + 10 + 8
		case 19:   return countOnes(hundredsPlace) + 10 + 8
		default:   break
	}
	switch x {
		case 10:   return 3
		case 100:  return 10
		case 1000: return 11
		default:   break
	}
	
	
	
	letterCount += countOnes(onesPlace)
	if x > 10 {
		letterCount += countTens(tensPlace )
	}
	if x > 100 {
		if x % 100 == 0 {
			letterCount += countOnes(hundredsPlace) + 7
		} else {
			letterCount += countOnes(hundredsPlace) + 10 // hundred and
		}
	}
	//fmt.Println(x, " ", onesPlace, " ", tensPlace, " ", hundredsPlace, letterCount)
	return
}
func countOnes(x int) (letterCount int) {
	switch x {
	case 0: return 0
	case 1: return 3
	case 2: return 3
	case 3: return 5
	case 4: return 4
	case 5: return 4
	case 6: return 3
	case 7: return 5
	case 8: return 5
	case 9: return 4
	}
	return
}
func countTens(x int) (letterCount int) {
	switch x {
	case 0: return 0
	case 1: return 3
	case 2: return 6
	case 3: return 6
	case 4: return 5
	case 5: return 5
	case 6: return 5
	case 7: return 7
	case 8: return 6
	case 9: return 6
	}
	return
}
func countHundreds(x int) (letterCount int) {
	return
}