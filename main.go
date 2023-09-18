package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main(){
	//var variableName string="Hello world";
	// variableName:=getVariableName()
	// fmt.Println(variableName)
	//slice()
	//cards:=newDeck()
	// hand,remainingCards:=deal(cards,5)
	// hand.print()
	// remainingCards.print()
	// cards.toString()
	// cards.saveToFile("deck")
	// cards := newDeckFromFile("deck")
	// cards.shuffle()
	//cards.print()
	bob:= person{
		firstName: "bob",
		lastName: "willy",
	}
	fmt.Printf("%+v",bob)
}

func getVariableName() string {
	return "Hello world"
}