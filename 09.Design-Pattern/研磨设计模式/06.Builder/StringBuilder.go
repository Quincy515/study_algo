package Builder

type StringBuilder struct {
	result string
}

func (b *StringBuilder) Part1() {
	b.result += "1"
}

func (b *StringBuilder) Part2() {
	b.result += "2"
}

func (b *StringBuilder) Part3() {
	b.result += "3"
}

func (b *StringBuilder) GetResult() string {
	return b.result
}
