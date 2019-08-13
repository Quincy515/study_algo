package HashTable

import (
	"fmt"
	"github.com/custergo/study_algo/Play-with-Data-Structures/14-Hash-Table/06-Resizing-in-Hash-Table/RBTree"
	"hash/fnv"
	"strconv"
)

const upperTol = 10    // 平均每个地址哈希冲突能容忍的上界
const lowerTol = 2     // 平均每个地址哈希冲突能容忍的下界
const initCapacity = 7 // 初始哈希表的容量

type HashTable struct {
	hashtable []*RBTree.RBTree
	M         int
	size      int
}

func NewHashTable(M int) *HashTable {
	var hashtable []*RBTree.RBTree
	for i := 0; i < M; i++ {
		hashtable = append(hashtable, RBTree.NewRBTree())
	}
	return &HashTable{hashtable, M, 0}
}

func (ht *HashTable) hash(key interface{}) int {
	return (int(HashCode(key)) & 0x7fffffff) % ht.M
}

func (ht *HashTable) GetSize() int {
	return ht.size
}

func (ht *HashTable) Add(key, value interface{}) {
	m := ht.hashtable[ht.hash(key)]
	if m.Contains(key) {
		m.Add(key, value)
	} else {
		m.Add(key, value)
		ht.size++

		if ht.size >= upperTol*ht.M {
			ht.resize(2 * ht.M)
		}
	}
}

func (ht *HashTable) Remove(key interface{}) interface{} {
	m := ht.hashtable[ht.hash(key)]
	if m.Contains(key) {
		ret := m.Remove(key)
		ht.size--

		if ht.size < lowerTol*ht.M && ht.M/2 > initCapacity {
			ht.resize(ht.M / 2)
		}
		return ret
	} else {
		return nil
	}
}

func (ht *HashTable) Set(key, value interface{}) {
	m := ht.hashtable[ht.hash(key)]
	if !m.Contains(key) {
		panic(fmt.Sprintf("%s doesn't exist!", key))
	}
	m.Set(key, value)
}

func (ht *HashTable) Contains(key interface{}) bool {
	return ht.hashtable[ht.hash(key)].Contains(key)
}

func (ht *HashTable) Get(key interface{}) interface{} {
	return ht.hashtable[ht.hash(key)].Get(key)
}

func (ht HashTable) resize(newM int) {
	var newHashTable []*RBTree.RBTree
	for i := 0; i < newM; i++ {
		newHashTable = append(newHashTable, RBTree.NewRBTree())
	}
	oldM := ht.M
	ht.M = newM
	for i := 0; i < oldM; i++ {
		m := ht.hashtable[i]
		for _, key := range m.KeySet() {
			newHashTable[ht.hash(key)].Add(key, m.Get(key))
		}
	}
	ht.hashtable = newHashTable
}

func HashCode(source interface{}) uint64 {
	switch source.(type) {
	case int:
		source = strconv.Itoa(source.(int))
	case float64:
		source = strconv.FormatFloat(source.(float64), 'f', 6, 64)
	}
	h := fnv.New64()
	h.Write([]byte(source.(string)))
	return h.Sum64()
}
