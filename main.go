package main

import (
	"github.com/mikako-shirai/Typeahead-Suggestion/data"
	"github.com/mikako-shirai/Typeahead-Suggestion/server"
)

func main() {
  
  // TODO 4:
  // 1. add toString method to stringify a Trie (to serialize)
  // 2. write a small function which runs every 1 minute and takes newTrie,
  // turns it into a string, and writes it to a file
  // 3. Optionally update TODO 3 to read from this file on startup, instead of building the Trie from a
  // text file of movie titles

  data.InitializeData()
  server.InitializeServer()
}
