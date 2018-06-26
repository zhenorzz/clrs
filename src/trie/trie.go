package trie

type RuneTrie struct {
	value    interface{}
	children map[rune]*RuneTrie
}

func New() *RuneTrie {
	return &RuneTrie{
		children: make(map[rune]*RuneTrie),
	}
}

func (trie *RuneTrie) Get(key string) interface{} {
	node := trie
	for _, r := range key {
		node = node.children[r]
		if node == nil {
			return nil
		}
	}
	return node.value
}

func (trie *RuneTrie) Put(key string, value interface{}) bool {
	node := trie
	for _, r := range key {
		child, ok := node.children[r]
		if ok == false {
			child = New()
			node.children[r] = child
		}
		node = child
	}

	isNewVal := node.value == nil
	node.value = value
	return isNewVal
}

// RuneTrie node and the rune key of the child the path descends into.
type nodeRune struct {
	node *RuneTrie
	r    rune
}

func (trie *RuneTrie) isLeaf() bool {
	return len(trie.children) == 0
}

func (trie *RuneTrie) Delete(key string) bool {
	path := make([]nodeRune, len(key))
	node := trie
	for i, r := range key {
		path[i] = nodeRune{node, r}
		node = node.children[r]
		if node == nil {
			return false
		}
	}
	node.value = nil
	if node.isLeaf() {
		for i := len(key) - 1; i >= 0; i-- {
			parent := path[i].node
			r := path[i].r
			delete(parent.children, r)
			if parent.value != nil || parent.isLeaf() == false {
				break
			}
		}
	}
	return true
}
