package Decorator

type AddDecorator struct {
	Component
	num int
}

func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func (add *AddDecorator) Calc() int {
	return add.Component.Calc() + add.num
}
