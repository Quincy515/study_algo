package Decorator

type MulDecorator struct {
	Component
	num int
}

func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (mul *MulDecorator) Calc() int {
	return mul.Component.Calc() * mul.num
}
