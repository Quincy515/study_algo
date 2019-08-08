package Trie

type node struct {
	isWord bool
	next   [26]*node
}

type trie struct {
	root *node
	size int
}

func generateNode() *node {
	return &node{next: [26]*node{}}
}

func Constructor() *trie {
	return &trie{root: generateNode()}
}

// 获得Trie中存储的单词数量
func (this *trie) GetSize() int {
	return this.size
}

// 向trie中添加一个新的单词word
func (this *trie) Add(word string) {
	cur := this.root

	for _, w := range word {
		offset := w - 'a'
		if offset > 26 {
			offset -= 26
		}

		if cur.next[offset] == nil {
			cur.next[offset] = generateNode()
		}
		cur = cur.next[offset]
	}

	if cur.isWord == false {
		cur.isWord = true
		this.size++
	}
}

// 查询单词word是否在trie中
func (this *trie) Contains(word string) bool {
	cur := this.root

	for _, w := range []rune(word) {
		offset := w - 'a'
		if cur.next[offset] == nil {
			return false
		}
		cur = cur.next[offset]
	}

	return cur.isWord
}
