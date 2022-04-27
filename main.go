package main

func main() {
	//cards := newdeck()
	// the above line of code doesnt modify the our previous slice but create a newone
	//fmt.Println(cards.toString())
	//cards.saveTofile("my_shit")
	//hand, remainingCards := deal(cards, 40)
	//hand.printfart()
	//remainingCards.printfart()

	// cards := readFile("my_shit")
	// cards.printfart()
	cards := newdeck()
	cards.shuffle()
	cards.printfart()
}
