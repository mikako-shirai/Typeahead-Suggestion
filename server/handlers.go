package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mikako-shirai/Typeahead-Suggestion/data"
)

var URL_GCP = "https://typeahead-suggestion.an.r.appspot.com/"

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", URL_GCP)
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

    body, err := ioutil.ReadAll(req.Body)
    check(err)
    word := string(body)

    data.NewTrie.Insert(word)
    fmt.Fprintf(w, "%v", string(word))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
