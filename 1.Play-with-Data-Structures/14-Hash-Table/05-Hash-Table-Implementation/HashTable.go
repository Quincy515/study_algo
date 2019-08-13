package HashTable

import (
	"fmt"
	"github.com/custergo/study_algo/Play-with-Data-Structures/14-Hash-Table/05-Hash-Table-Implementation/RBTree"
	"hash/fnv"
	"strconv"
)

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
	}
}

func (ht *HashTable) Remove(key interface{}) interface{} {
	m := ht.hashtable[ht.hash(key)]
	if m.Contains(key) {
		ret := m.Remove(key)
		ht.size--
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
