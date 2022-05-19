package main

import (
	"github.com/mikako-shirai/Typeahead-Suggestion/data"
	"github.com/mikako-shirai/Typeahead-Suggestion/server"
)

func main() {

    // TODO
    // - add toString method to stringify a trie (serialize trie)
    // - write a function which runs every 1-5 minute (takes new trie, turns it into a string, writes it to a file)
    // - maybe update seeding part to read from this^ file on startup instead of building the trie from the seeding data everytime?

    data.InitializeData()
    server.InitializeServer()
}
