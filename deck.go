package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func newDeck() deck{
	cards:=deck{}
	cardSuits:=deck{"Spades","diamonds","hearts","clubs"}
	cardValues:=deck{"Ace","Two","Three","Four"}

	for _,suit:=range cardSuits{
		for _,value:=range cardValues{
			cards=append(cards,value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck,deck){
	return d[:handSize],d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) print(){
	for i,card:=range d{
		fmt.Println(i, card)
	}
}