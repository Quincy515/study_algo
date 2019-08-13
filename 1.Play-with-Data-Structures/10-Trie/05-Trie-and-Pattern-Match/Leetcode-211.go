package LeetCode_211

type Node struct {
	isWord bool
	next   map[string]*Node
}

type WordDictionary struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{root: &Node{next: make(map[string]*Node)}}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for _, w := range []rune(word) {
		c := string(w)
		if cur.next[c] == nil {
			cur.next[c] = &Node{next: make(map[string]*Node)}
		}
		cur = cur.next[c]
	}
	cur.isWord = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.match(this.root, word, 0)
}

func (this *WordDictionary) match(n *Node, word string, index int) bool {
	if index == len(word) {
		return n.isWord
	}
	c := string([]rune(word)[index])
	if c != "." {
		if n.next[c] == nil {
			return false
		}
		return this.match(n.next[c], word, index+1)
	} else {
		for w := range n.next {
			if this.match(n.next[w], word, index+1) {
				return true
			}
		}
		return false
	}
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
