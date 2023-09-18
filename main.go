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
	// bob:= person{
	// 	firstName: "bob",
	// 	lastName: "willy",
	// 	contact: contactInfo{
	// 		email: "bob@gmail.com",
	// 		zipCode: "121323",
	// 	},
	// }
	// bob.updateName("joe")
	// fmt.Printf("%+v",bob)
	m:=make(map[string]string)
	m["color1"]="#ffff"
	m["color2"]="#eb12"
	// delete(m,"color1")
	printMap(m)
	fmt.Println(m)
}

func printMap(m map[string]string){
	for key,value:= range m{
		fmt.Println("Key is", key, "Value is", value)
	}
}

func (personPointer *person) updateName(name string){
	(*personPointer).firstName = name;
}

func getVariableName() string {
	return "Hello world"
}