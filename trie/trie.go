package trie

const ALPHABET_ALL_LENGTH = 26
var ALPHABET_ALL = []string{
    "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
    "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
}

type Trie struct {
    root *TrieNode
}

type TrieNode struct {
    children [ALPHABET_ALL_LENGTH]*TrieNode
    isWord bool
}

func InitializeTrie() *Trie {
    return &Trie{
        root: &TrieNode{},
    }
}

func (trie *Trie) Insert(word string) {
    length := len(word)
    currentNode := trie.root

    for i := 0; i < length; i++ {
        index := word[i] - 'a'
        if currentNode.children[index] == nil {
            currentNode.children[index] = &TrieNode{}
        }
        currentNode = currentNode.children[index]
    }
    currentNode.isWord = true
}

func (trie *Trie) SearchByWord(word string) bool {
    length := len(word)
    currentNode := trie.root
  
    for i := 0; i < length; i++ {
        index := word[i] - 'a'
        if currentNode.children[index] == nil {
            return false
        }
        currentNode = currentNode.children[index]
    }
    return currentNode.isWord
}

func (trie *Trie) SearchByPrefix(prefix string) []string {
    foundWords := []string{}

    if isValidPrefix := trie.SearchByWord(prefix); isValidPrefix {
        foundWords = append(foundWords, prefix)
    }

    nodeAtPrefix := trie.getNodeByWord(prefix)
    foundWords = nodeAtPrefix.getWordsByPrefix(prefix, foundWords)
    
    return foundWords
}

func (trie *Trie) getNodeByWord(word string) *TrieNode {
    length := len(word)
    currentNode := trie.root
  
    for i := 0; i < length; i++ {
        index := word[i] - 'a'
        currentNode = currentNode.children[index]
    }
    return currentNode
}

func (node *TrieNode) getWordsByPrefix(prefix string, foundWords []string) []string {
    for index, node := range &node.children {
        if node != nil {
            char := ALPHABET_ALL[index]
            if node.isWord {
                foundWords = append(foundWords, prefix + char)
            }
            node.getWordsByPrefix(prefix + char, foundWords)
        }
    }
    return foundWords
}

var NewTrie = InitializeTrie()
