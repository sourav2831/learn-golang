package main

import "fmt"

type contactInfo struct {
	email string
	zipCode string
}
type person struct {
	firstName string
	lastName string
	contact contactInfo
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
		contact: contactInfo{
			email: "bob@gmail.com",
			zipCode: "121323",
		},
	}
	fmt.Printf("%+v",bob)
}

func getVariableName() string {
	return "Hello world"
}