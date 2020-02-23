package Observer

import "fmt"

// Reader: 读取消息类
type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{name: name}
}

func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}
