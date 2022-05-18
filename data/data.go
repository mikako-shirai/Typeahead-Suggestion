package data

import (
	"fmt"

	"github.com/mikako-shirai/Typeahead-Suggestion/trie"
)

// TODO 3:
// Replace these strings with contents of a file
// 1. find a file of movie titles
// 2. read a file and add each string to NewTrie

var words = []string{
    "sam", "john", "tim", "jose", "rose", "roses",
    "rosen", "rosetta", "cat", "dog", "dogg", "roses",
}

var wordsToFind = []string{
    "sam", "john", "tim", "jose", "rose", "cat",
    "dog", "dogg", "roses", "rosess", "ans", "san",
}

func seedData() {
    for i := 0; i < len(words); i++ {
        trie.NewTrie.Insert(words[i])
    }
}

func checkWords() {
    for i := 0; i < len(wordsToFind); i++ {
        found := trie.NewTrie.SearchByWord(wordsToFind[i])
        if !found {
            fmt.Printf("[init] \"%s\" not found\n", wordsToFind[i])
        }
    }
}

func InitializeData() {
    seedData()
    checkWords()
}
