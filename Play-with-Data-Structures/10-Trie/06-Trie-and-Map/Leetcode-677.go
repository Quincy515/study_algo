package LeetCode_677

type Node struct {
	value int
	next  map[string]*Node
}

type MapSum struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{root: &Node{next: make(map[string]*Node)}}
}

func (this *MapSum) Insert(key string, val int) {
	cur := this.root
	for _, w := range []rune(key) {
		c := string(w)
		if cur.next[c] == nil {
			cur.next[c] = &Node{next: make(map[string]*Node)}
		}
		cur = cur.next[c]
	}
	cur.value = val
}

func (this *MapSum) Sum(prefix string) int {
	cur := this.root
	for _, w := range []rune(prefix) {
		c := string(w)
		if cur.next[c] == nil {
			return 0
		}
		cur = cur.next[c]
	}
	return this.sum(cur)
}

func (this *MapSum) sum(n *Node) int {
	//if len(n.next) == 0 {
	//	return n.value
	//} 递归终止条件包含在了res中

	res := n.value
	for s := range n.next {
		res += this.sum(n.next[s])
	}
	return res
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
