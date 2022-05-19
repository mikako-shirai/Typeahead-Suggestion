package data

import (
	"github.com/mikako-shirai/Typeahead-Suggestion/trie"
)

var NewTrie = trie.InitializeTrie()

func seedData() {
	dictWords := readDataFromFile()
	for i := 0; i < len(dictWords); i += 200 {
		NewTrie.Insert(dictWords[i])
	}
}

func InitializeData() {
    seedData()
}
