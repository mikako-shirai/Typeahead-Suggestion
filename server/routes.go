package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mikako-shirai/Typeahead-Suggestion/trie"
)

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}

func getAllWords(w http.ResponseWriter, req *http.Request) {
    enableCors(&w)
    
    foundWords := trie.NewTrie.GetAllWords()
    words, err := json.Marshal(foundWords)

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Fprintf(w, "%v", string(words))
    }
}

func getWordsByPrefix(w http.ResponseWriter, req *http.Request) {
    enableCors(&w)

    currentWord := req.URL.Query().Get("word")
    currentWord = strings.ToLower(currentWord)

    foundWords := trie.NewTrie.GetWordsByPrefix(currentWord)
    words, err := json.Marshal(foundWords)

    if err != nil {
        fmt.Println(err)
    } else {
        if len(foundWords) == 0 {
            trie.NewTrie.Insert(currentWord)
        }
        fmt.Fprintf(w, "%v", string(words))
    }
}

func insert(w http.ResponseWriter, req *http.Request) {
    enableCors(&w)

    currentWord := req.URL.Query().Get("word")
    currentWord = strings.ToLower(currentWord)
    trie.NewTrie.Insert(currentWord)
}
