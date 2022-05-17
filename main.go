package main

import (
	"fmt"
	"net/http"

	"github.com/mikako-shirai/Typeahead-Suggestion/server"
)

// TODO 3:
// Replace these strings with contents of a file
// 1. find a file of movie titles
// 2. read a file and add each string to NewTrie

var words = []string{"sam", "john", "tim", "jose", "rose",
  "cat", "dog", "dogg", "roses"}

var wordsToFind = []string{"sam", "john", "tim", "jose", "rose",
  "cat", "dog", "dogg", "roses", "rosess", "ans", "san"}

func main() {
  for i := 0; i < len(words); i++ {
    server.NewTrie.Insert(words[i])
  }

  for i := 0; i < len(wordsToFind); i++ {
    found := server.NewTrie.IsInTrie(wordsToFind[i])
    if found {
      fmt.Printf("Word \"%s\" found in Trie\n", wordsToFind[i])
    } else {
      fmt.Printf("Word \"%s\" not found in Trie\n", wordsToFind[i])
    }
  }

  // TODO 4:
  // 1. add toString method to stringify a Trie (to serialize)
  // 2. write a small function which runs every 1 minute and takes newTrie,
  // turns it into a string, and writes it to a file
  // 3. Optionally update TODO 3 to read from this file on startup, instead of building the Trie from a
  // text file of movie titles

  http.HandleFunc("/headers", server.Headers)
  http.HandleFunc("/search", server.Search)

  http.ListenAndServe(":8090", nil)
}
