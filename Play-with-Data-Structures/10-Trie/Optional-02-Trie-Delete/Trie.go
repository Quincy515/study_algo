package Trie

type Node struct {
	isWord bool
	next   map[string]*Node
}

type Trie struct {
	root *Node
	size int
}

func NewNode() *Node {
	return &Node{next: make(map[string]*Node)}
}

func NewTrie() *Trie {
	return &Trie{root: NewNode()}
}

// 获取Trie中存储的单词数量
func (t *Trie) GetSize() int {
	return t.size
}

// 向Trie中添加一个新的单词word
func (t *Trie) Add(word string) {
	cur := t.root
	for _, w := range []rune(word) {
		c := string(w)
		if cur.next[c] == nil {
			cur.next[c] = NewNode()
		}
		cur = cur.next[c]
	}

	if cur.isWord == false {
		cur.isWord = true
		t.size++
	}
}

// 查询单词word是否在Trie中
func (t *Trie) Contains(word string) bool {
	cur := t.root
	for _, w := range []rune(word) {
		c := string(w)
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}
	return cur.isWord
}

// 查询是否在Trie中有单词以prefix为前缀
func (t *Trie) IsPrefix(prefix string) bool {
	cur := t.root
	for _, s := range []rune(prefix) {
		c := string(s)
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}
	return true
}

// 删除word，返回是否删除成功
func (t *Trie) Remove(word string) bool {
	// 将搜索沿路的节点放入栈中
	stack := make([]*Node, 0)
	stack = append(stack, t.root)
	cur := stack[len(stack)-1]

	for i := 0; i < len(word); i++ {
		c := string([]rune(word)[i])
		if cur.next[c] == nil {
			return false
		}
		stack = append(stack, cur.next[c])
	}

	if !cur.isWord {
		return false
	}

	// 将该单词结尾isWord置空
	cur.isWord = false
	t.size--

	// 如果单词最后一个字母的节点的next非空
	// 说明trie中还存储了其他以该单词为前缀的单词，直接返回
	if len(cur.next) > 0 {
		return true
	} else {
		stack = stack[:len(stack)-1]
	}

	// 自底向上删除
	for i := len(word) - 1; i >= 0; i-- {
		//c := string([]rune(word)[i])
		//cur.next[c]
		// 如果一个节点的isWord为true，或者是其他单词的前缀，则直接返回
		if cur.isWord || len(cur.next) > 0 {
			return true
		}
		stack = stack[:len(stack)-1]
	}
	return true
}
