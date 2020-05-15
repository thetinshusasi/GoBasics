package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func (d deck) Print() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, numOfCards int) (deck, deck) {
	return d[:numOfCards], d[numOfCards:]
}
func (d deck) toString() string {

	dList := []string(d)
	msg := strings.Join(dList, ",")

	return msg
}

func deckStringToDec(deckString string) deck {
	return deck(strings.Split(deckString, ","))
}

func (d deck) saveToFile(fileName string) error {
	err := ioutil.WriteFile(fileName, []byte(d.toString()), 7777)
	return err
}

func readDeckFile(fileName string) (deck, error) {
	dBytes, err := ioutil.ReadFile(fileName)
	dataDeck := deckStringToDec(string(dBytes))
	return dataDeck, err
}
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d { // i = index(go naming convention)
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i] //swap the elements
	}
}
