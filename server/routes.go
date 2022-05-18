package server

import (
	"fmt"
	"net/http"

	"github.com/mikako-shirai/Typeahead-Suggestion/trie"
)

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func search(w http.ResponseWriter, req *http.Request) {
    paramWord := req.URL.Query().Get("word")
    fmt.Printf("[GET] param : \"%s\"\n", paramWord)

    // TODO 2:
    // Use Prefix method instead of IsInTree

    found := trie.NewTrie.SearchByWord(paramWord)
    response := make(map[string]string)

    if found {
        response["found"] = "found"
        response["word"] = paramWord
        fmt.Printf("[GET] \"%s\" found : TRUE\n", paramWord)
        fmt.Fprintf(w, "%v\n", response)
    } else {
        response["found"] = "not found"
        fmt.Printf("[GET] \"%s\" found : false\n", paramWord)
        trie.NewTrie.Insert(paramWord)
        fmt.Fprintf(w, "%v\n", response)
    }
}
