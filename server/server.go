package server

import (
	"fmt"
	"net/http"

	"github.com/mikako-shirai/Typeahead-Suggestion/trie"
)

var NewTrie = trie.InitializeTrie()

func Search(w http.ResponseWriter, req *http.Request) {
  paramWord := req.URL.Query().Get("word")
  fmt.Println("GET param :", paramWord)

  // TODO 2:
	// Use Prefix method instead of IsInTree
  found := NewTrie.IsInTrie(paramWord)
  if found {
    fmt.Printf("Word \"%s\" : found\n", paramWord)
  } else {
    fmt.Printf("Word \"%s\" : not found\n", paramWord)
    NewTrie.Insert(paramWord)
  }

  fmt.Fprintf(w, "search\n")
}

func Headers(w http.ResponseWriter, req *http.Request) {
  for name, headers := range req.Header {
    for _, h := range headers {
      fmt.Fprintf(w, "%v: %v\n", name, h)
    }
  }
}
