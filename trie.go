package main

type TrieNode struct {
	Children map[rune]*TrieNode
	IsWord   bool
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: &TrieNode{
			Children: make(map[rune]*TrieNode),
		},
	}
}

func (t *Trie) Insert(word string) {
	node := t.Root
	for _, ch := range word {
		if _, exists := node.Children[ch]; !exists {
			node.Children[ch] = &TrieNode{
				Children: make(map[rune]*TrieNode),
			}
		}
		node = node.Children[ch]
	}
	node.IsWord = true
}

func (t *Trie) findNode(prefix string) *TrieNode {
	node := t.Root
	for _, ch := range prefix {
		next, exists := node.Children[ch]
		if !exists {
			return nil
		}
		node = next
	}
	return node
}

func (t *Trie) collect(
	node *TrieNode,
	current string,
	results *[]string,
) {
	if node.IsWord {
		*results = append(*results, current)
	}
	for ch, child := range node.Children {
		t.collect(
			child,
			current+string(ch),
			results,
		)
	}
}

func (t *Trie) AutoComplete(prefix string) []string {
	node := t.findNode(prefix)
	if node == nil {
		return nil
	}
	var results []string
	t.collect(
		node,
		prefix,
		&results,
	)
	return results
}

