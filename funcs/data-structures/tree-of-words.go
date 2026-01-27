package datastructures

import "log"

func TreeMain() {
	res := Trie([]string{"bimba", "unitaz", "unita"}, [][]string{
		{"search", "unitaz"},
		{"delete", "unitaz"},
		{"search", "unitaz"},
		{"delete", "unitaz"},
		{"delete", "unito"},
		{"search", "uni"},
		{"search", "unit"},
		{"search", "unita"},
		{"search", "bimba"},
	})

	log.Println(res)
}

type TrieNode struct {
	Children map[rune]*TrieNode
	End      bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[rune]*TrieNode),
		End:      false,
	}
}

func createTrie(words []string) *TrieNode {
	// === DO NOT MODIFY ===
	root := NewTrieNode()
	for _, word := range words {
		insert(root, word)
	}
	return root
}

func insert(root *TrieNode, word string) {
	// === DO NOT MODIFY ===
	node := root
	for _, char := range word {
		if _, exists := node.Children[char]; !exists {
			node.Children[char] = NewTrieNode()
		}
		node = node.Children[char]
	}
	node.End = true
}

func search(root *TrieNode, word string) bool {
	node := root
	for _, char := range word {
		nextNode, ok := node.Children[char]
		if !ok {
			return false
		}
		node = nextNode
	}

	return node.End
}

func deleteWord(root *TrieNode, word string) bool {
	node := root
	for _, char := range word {
		nextNode, ok := node.Children[char]
		if !ok {
			return false
		}
		node = nextNode
	}

	if node.End {
		node.End = false
		return true
	}

	return false
}

func Trie(initialWords []string, commands [][]string) []bool {
	// === DO NOT MODIFY ===
	root := createTrie(initialWords)

	var output []bool
	for _, command := range commands {
		if command[0] == "search" {
			output = append(output, search(root, command[1]))
		} else if command[0] == "delete" {
			output = append(output, deleteWord(root, command[1]))
		}
	}
	return output
}
