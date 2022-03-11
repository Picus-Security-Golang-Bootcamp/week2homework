package main

import (
	"fmt"
	"sort"
	"strings"
)

type list struct {
	book string
}

func main() {
	library := []list{{"gameofthrones"}, {"olivertwist"}, {"harrypotter"}, {"lotr"}}

	sort.Slice(library, func(i, j int) bool {

		return library[i].book <= library[j].book
	})

	var searchingBook string

	fmt.Scanf("%s", &searchingBook)
	fmt.Printf("Searching book: %v\n", searchingBook)

	searchingBook = strings.ToLower(searchingBook)

	x := sort.Search(len(library), func(i int) bool {

		return string(library[i].book) >= searchingBook

	})

	if x < len(library) && library[x].book == searchingBook {
		fmt.Println("Found on list")
	} else {
		fmt.Println("Found nothing")
	}
}
