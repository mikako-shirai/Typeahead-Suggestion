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

    check(err)
    fmt.Fprintf(w, "%v", string(words))
}

func getWordsByPrefix(w http.ResponseWriter, req *http.Request) {
    enableCors(&w)

    prefix := req.URL.Query().Get("prefix")
    foundWords := data.NewTrie.GetWordsByPrefix(prefix)
    words, err := json.Marshal(foundWords)

    check(err)
    fmt.Fprintf(w, "%v", string(words))
}

func insert(w http.ResponseWriter, req *http.Request) {
    enableCors(&w)

    word := req.URL.Query().Get("word")
    fmt.Printf("received word : '%v'\n", word)
    data.NewTrie.Insert(word)
    fmt.Fprintf(w, "%v", string(word))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
