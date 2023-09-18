package main

func main(){
	//var variableName string="Hello world";
	// variableName:=getVariableName()
	// fmt.Println(variableName)
	//slice()
	cards:=newDeck()
	// hand,remainingCards:=deal(cards,5)
	// hand.print()
	// remainingCards.print()
	// cards.toString()
	// cards.saveToFile("deck")
	// cards := newDeckFromFile("deck")
	cards.shuffle()
	cards.print()
}

func getVariableName() string {
	return "Hello world"
}