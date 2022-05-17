package trie

const ALBHABET_ALL_LENGTH = 26

type Trie struct {
  root *TrieNode
}

type TrieNode struct {
  children [ALBHABET_ALL_LENGTH]*TrieNode
  isWord bool
}

func InitializeTrie() *Trie {
  return &Trie{
    root: &TrieNode{},
  }
}

func (trie *Trie) Insert(word string) {
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

func (trie *Trie) IsInTrie(word string) bool {
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

// TODO 1:
// Add Prefix method for returning subtrees of a search
