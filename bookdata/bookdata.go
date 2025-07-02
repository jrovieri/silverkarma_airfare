package main

import (
	"fmt"
	"strings"
)

var books = []string{
	"The Tale of Two Cities",
	"The Joy of Compassion",
	"From Earth to Sirius",
	"Galaxy of Three Pebbles",
	"House of Impatience",
	"Raffle Got the Blues",
	"Where the Rainbows Have No Ends",
	"Storm, Typhoons, Snow, and Sunshine",
	"You Can't Say That Ever Again",
	"Tally of the Seashells",
	"For Whom the Snails Toast",
	"Where Did All the Greek Gods Go",
	"The Elemental Drift",
	"Earthers Beware",
	"Welcome to Our Giant House on the Hudson",
	"I Have Something Not Nice to Say",
	"The Dinner Table of Awkwardness",
	"Sunshine is Overrated",
	"The Planet of Green Cities",
	"Remembering the Yellow Train Tracks",
	"Seasick on the Land",
	"Blue Infini",
	"Blue Planet Red Ambition",
}

func main() {

	var searchString string

	fmt.Print("Enter a string to search in book titles: ")
	fmt.Scanln(&searchString)

	for _, title := range books {

		if strings.Contains(strings.ToLower(title), strings.ToLower(searchString)) {
			fmt.Println(title)
		}
	}
}
