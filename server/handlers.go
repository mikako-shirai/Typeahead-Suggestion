package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mikako-shirai/Typeahead-Suggestion/data"
)

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}

func getAllWords(w http.ResponseWriter, req *http.Request) {
    enableCors(&w)

    foundWords := data.NewTrie.GetAllWords()
    words, err := json.Marshal(foundWords)

    if err != nil {
        panic(err)
    } else {
        fmt.Fprintf(w, "%v", string(words))
    }
}

func getWordsByPrefix(w http.ResponseWriter, req *http.Request) {
    enableCors(&w)

    prefix := req.URL.Query().Get("word")
    foundWords := data.NewTrie.GetWordsByPrefix(prefix)
    words, err := json.Marshal(foundWords)

    if err != nil {
        panic(err)
    } else {
        fmt.Fprintf(w, "%v", string(words))
    }
}

func insert(w http.ResponseWriter, req *http.Request) {
    enableCors(&w)

    word := req.URL.Query().Get("word")
    data.NewTrie.Insert(word)
}
