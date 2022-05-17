package main

import (
	"fmt"
	"net/http"
)

const ALBHABET_ALL = 26

type TrieNode struct {
	children [ALBHABET_ALL]*TrieNode
	isWord bool
}

type Trie struct {
	root *TrieNode
}

func initTrie() *Trie {
	return &Trie{
		root: &TrieNode{},
	}
}

func (trie *Trie) insert(word string) {
	wordLength := len(word)
	currentNode := trie.root

	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if currentNode.children[index] == nil {
			currentNode.children[index] = &TrieNode{}
		}
		currentNode = currentNode.children[index]
	}
	currentNode.isWord = true
}

func (trie *Trie) isInTree(word string) bool {
	wordLength := len(word)
	currentNode := trie.root
	
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if currentNode.children[index] == nil {
			return false
		}
		currentNode = currentNode.children[index]
	}

	return currentNode.isWord
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func search(w http.ResponseWriter, req *http.Request) {
	paramWord := req.URL.Query().Get("word")
	fmt.Println("GET param :", paramWord)

	found := newTrie.isInTree(paramWord)
	if found {
		fmt.Printf("Word \"%s\" : found\n", paramWord)
	} else {
		fmt.Printf("Word \"%s\" : not found\n", paramWord)
		newTrie.insert(paramWord)
	}

	fmt.Fprintf(w, "search\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

var newTrie = initTrie()

func main() {
	words := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses"}
	wordsToFind := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses", "rosess", "ans", "san"}

	for i := 0; i < len(words); i++ {
		newTrie.insert(words[i])
	}

	for i := 0; i < len(wordsToFind); i++ {
		found := newTrie.isInTree(wordsToFind[i])
		if found {
			fmt.Printf("Word \"%s\" found in Trie\n", wordsToFind[i])
		} else {
			fmt.Printf("Word \"%s\" not found in Trie\n", wordsToFind[i])
		}
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/search", search)

	http.ListenAndServe(":8090", nil)
}