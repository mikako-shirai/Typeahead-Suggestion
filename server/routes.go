package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mikako-shirai/Typeahead-Suggestion/trie"
)

func search(w http.ResponseWriter, req *http.Request) {
    currentWord := req.URL.Query().Get("word")
    currentWord = strings.ToLower(currentWord)

    foundWords := trie.NewTrie.SearchByPrefix(currentWord)
    fmt.Printf("foundWords  %v\n\n", foundWords)

    // words, err := json.Marshal(foundWords)
    // if err != nil {
    //     fmt.Println(err)
    // }
    
    if len(foundWords) != 0 {
        fmt.Printf("[GET] \"%s\" found\n", currentWord)
        fmt.Fprintf(w, "%v\n", foundWords)
    } else {
        trie.NewTrie.Insert(currentWord)
        
        fmt.Printf("[GET] \"%s\" not found\n", currentWord)
        fmt.Fprintf(w, "%v\n", foundWords)
    }
}

func insert(w http.ResponseWriter, req *http.Request) {
    currentWord := req.URL.Query().Get("word")
    currentWord = strings.ToLower(currentWord)
    trie.NewTrie.Insert(currentWord)
}
