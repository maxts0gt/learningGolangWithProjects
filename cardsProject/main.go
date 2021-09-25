package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
//	cards := newDeck()
//	err := cards.saveToFile("my_cards")
//	if err != nil {
//		return
//	}
//}
