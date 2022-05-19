package data

import (
	"github.com/mikako-shirai/Typeahead-Suggestion/trie"
)

var NewTrie = trie.InitializeTrie()

func seedData() {
    for i := 0; i < len(words); i++ {
        NewTrie.Insert(words[i])
    }
}

// func checkWords() {
//     for i := 0; i < len(wordsToFind); i++ {
//         found := trie.NewTrie.SearchByWord(wordsToFind[i])
//         if !found {
//             fmt.Printf("[init] \"%s\" not found\n", wordsToFind[i])
//         }
//     }
// }

func InitializeData() {
    seedData()
    // checkWords()
}
