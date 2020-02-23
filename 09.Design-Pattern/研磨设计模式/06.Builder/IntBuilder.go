package Builder

type IntBuilder struct {
	result int64
}

func (b *IntBuilder) Part1() {
	b.result += 1
}

func (b *IntBuilder) Part2() {
	b.result += 2
}

func (b *IntBuilder) Part3() {
	b.result += 3
}

func (b *IntBuilder) GetResult() int64 {
	return b.result
}
