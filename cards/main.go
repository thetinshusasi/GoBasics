package main

import (
	"fmt"
	"os"
)

var card = "tos"

func main() {
	cards := newDeck()
	// cards.Print()
	hand, cards := deal(cards, 5)
	hand.saveToFile("testFile.txt")
	newDeck, err := readDeckFile("testFile1.txt")
	if err == nil {

		newDeck.Print()
	} else {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	// cards.Print()

}
