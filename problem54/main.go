package main

import (
	"time"
	"sort"
	"bufio"
	"strings"
	"os"
	"fmt"
)

const path = "p054_poker.txt"

const (
	straightFlush = 80000000
	fourKind      = 70000000
	fullHouse     = 60000000
	flush         = 50000000
	straight      = 40000000
	threeKind     = 30000000
	twoPair       = 20000000
	onePair       = 10000000
)

//weights for the two pair method to ensure that there is no ambiquity based on high cards
const (
	highPair = 14 * 14 * 14 * 14
	lowPair  = 14 * 14
	
)

type card struct {
	rank int //2 - 14
	suit int //1 - 4
}
type hand struct {
	cards   []card
	handVal int
}
type game struct {
	h1 hand
	h2 hand
}
func findAnswer(path string) (ans int) {
	g := parseInput(path)
	for i := range g {
		v1 := g[i].h1.calculateValue()
		v2 := g[i].h2.calculateValue()
		if v1 > v2 {
			ans++
		}
	}
	return
}
func (h hand)calculateValue() (int){
	sort.Sort(h)

	//straight flush
	if hasStraightFlush(h) {
		return straightFlush + highCard(h)
	} 
	//4 of a kind
	flag,val := has4(h) 
	if flag {
		return fourKind + val
	}
	//full house + 3 rank
	flag, val = hasFullhouse(h)
	if flag {
		return fullHouse + val
	}
	// flush + high card
	if hasflush(h) {
		return flush + highCard(h)
	}
	//straight + high card
	if hasStraight(h) {
		return straight + highCard(h)
	}
	//3s + card
	flag, val = has3(h)
	if flag {
		return threeKind + val
	} 
	//2pair
	flag, val = hasTwoPair(h)
	if flag {
		return twoPair + val
	}
	//1 pair 
	flag, val = hasPair(h)
	if flag {
		return onePair + val
	}
	//high card
	return highCard(h)
}

//checks for a flush, returns true if all suits match
func hasflush(h hand) (bool){
	suit := h.cards[0].suit //first card's suit
	for i := range h.cards {
		if h.cards[i].suit != suit {
			return false //suit mismatch
		}
	}
	return true
}
//checks for a straight
func hasStraight(h hand)(bool) {
	rank := h.cards[0].rank
	//Ace can be in a straight two different ways
	//lazy fix
	if rank == 1 {
		return (
			(h.cards[1].rank == 2 && 
			 h.cards[2].rank == 3 && 
			 h.cards[3].rank == 4 && 
			 h.cards[4].rank == 5) ||
			(h.cards[1].rank == 10 && 
			 h.cards[2].rank == 11 && 
			 h.cards[3].rank == 12 && 
			 h.cards[4].rank == 13))
	}
	for i := 0; i < 5; i++ {
		if rank + i != h.cards[i].rank {
			return false
		}
	}
	return true
}

func hasStraightFlush(h hand) (bool){
	return (hasflush(h) && hasStraight(h))
}

//4 of a kind can only be in the form:
//a,a,a,a,b;b,a,a,a,a
//since the cards are sorted by rank
func has4(h hand) (bool,int){
	if h.cards[0].rank == h.cards[3].rank {
		return true, h.cards[0].rank
	}
	if h.cards[1].rank == h.cards[4].rank {
		return true, h.cards[1].rank
	}
	return false,-1
}
//3 of a kind forms
//a,a,a,b,c;b,a,a,a,c;b,c,a,a,a
func has3(h hand) (bool, int) {
	if h.cards[0].rank == h.cards[2].rank {
		return true, h.cards[0].rank
	}
	if h.cards[1].rank == h.cards[3].rank {
		return true, h.cards[1].rank
	}
	if h.cards[2].rank == h.cards[4].rank {
		return true, h.cards[2].rank
	}
	return false,-1
}
//full house forms
//a,a,b,b,b;b,b,b,a,a
func hasFullhouse(h hand) (bool,int){
	if (h.cards[0].rank == h.cards[1].rank && 
	h.cards[2].rank == h.cards[4].rank) {
		return true, h.cards[2].rank
	}
	if (h.cards[0].rank == h.cards[2].rank &&
	h.cards[3].rank == h.cards[4].rank) {
		return true, h.cards[0].rank
	}
	return false, -1
}
//two pair is a little bit spaghetti in order to accurately rank hands
//probably a more elegant way to do this
func hasTwoPair(h hand) (bool,int) {
	high := highCard(h)
	//a,a,c,b,b
	if (h.cards[0].rank == h.cards[1].rank && 
	    h.cards[3].rank == h.cards[4].rank ) {
		//ensure aces are counted as high pair
		if high == 14 {
			return true, highPair * h.cards[0].rank + lowPair * h.cards[3].rank + h.cards[2].rank
		} else {
			return true, highPair * h.cards[3].rank + lowPair * h.cards[0].rank + h.cards[2].rank
		}
	}
	//c,a,a,b,b - we don't need to do the extra check for this case, low pair can't be ace
	if (h.cards[1].rank == h.cards[2].rank &&
		h.cards[3].rank == h.cards[4].rank) {
		return true, highPair * h.cards[1].rank + lowPair * h.cards[3].rank + h.cards[0].rank
	}
	//a,a,b,b,c
	if (h.cards[0].rank == h.cards[1].rank && 
		h.cards[2].rank == h.cards[3].rank) {

		if high == 14 {
			return true, highPair * h.cards[0].rank + lowPair * h.cards[3].rank + h.cards[4].rank
		} else {
			return true, highPair * h.cards[2].rank + lowPair * h.cards[0].rank + h.cards[4].rank
		}
	}
	return false, -1
}

func hasPair(h hand) (bool,int){
	high := highCard(h)
	//a,a,b,c,d
	if h.cards[0].rank == h.cards[1].rank {
		return true, lowPair * cardVal(h.cards[0]) + cardVal(h.cards[4])
	}
	//b,a,a,c,d
	if h.cards[1].rank == h.cards[2].rank {
		if high == 14 {
			return true, lowPair * cardVal(h.cards[1]) + high
		}
		return true, lowPair * cardVal(h.cards[1]) + cardVal(h.cards[4])
	}
	//b,c,a,a,d
	if h.cards[2].rank == h.cards[3].rank {
		if high == 14 {
			return true, lowPair * cardVal(h.cards[2]) + high
		}
		return true, lowPair * cardVal(h.cards[2]) + cardVal(h.cards[4])
	}
	//b,c,d,a,a
	if h.cards[3].rank == h.cards[4].rank {
		if high == 14 {
			return true, lowPair * cardVal(h.cards[3]) + high
		}
		return true, lowPair * cardVal(h.cards[3]) + cardVal(h.cards[2])
	}

	return false, -1
}
func highCard(h hand) (int) {
	if h.cards[0].rank == 1 {
		return 14 //ace high
	}
	return h.cards[4].rank
}

func cardVal(c card) (int){
	if c.rank == 1 {
		return 14
	}
	return c.rank
}

//functions to implement sort interface
func (h hand)Len()(int) {
	return len(h.cards)
}
func (h hand)Less(i,j int)(bool) {
	//sort by rank first because order matters for straights but not flushes
	if h.cards[i].rank < h.cards[j].rank {
		return true
	} else if h.cards[i].rank > h.cards[j].rank {
		return false
	} 
	//if rank is equal, sort based on suit
	if h.cards[i].suit < h.cards[j].suit {
		return true
	} 
	return false
}
func (h hand)Swap(i,j int) {
	h.cards[i],h.cards[j] = h.cards[j],h.cards[i]
}


func parseInput(path string) (games []game) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		var tmp game
		for i := 0; i < 5; i++ {
			tmp.h1.cards = append(tmp.h1.cards,parseCard(line[i]))
			tmp.h2.cards = append(tmp.h2.cards,parseCard(line[i+5]))
		}
		games = append(games,tmp)
	}

	return
}

func parseCard(in string) (c card) {
	rank := in[0]
	suit := in[1]
	switch rank {
	case 'T': c.rank = 10 //ten
	case 'J': c.rank = 11 //jack
	case 'Q': c.rank = 12 //queen
	case 'K': c.rank = 13 //king
	case 'A': c.rank = 1  //ace
	default:  c.rank = int(rank - '0')
	}
	switch suit {
	case 'C': c.suit = 1 //club
	case 'H': c.suit = 2 //heart
	case 'S': c.suit = 3 //spade
	case 'D': c.suit = 4 //diamond
	}
	return
}

func main() {
	t0 := time.Now()
	ans := findAnswer(path)
	fmt.Println("Answer: ",ans,"\nTime: ",time.Since(t0))
}