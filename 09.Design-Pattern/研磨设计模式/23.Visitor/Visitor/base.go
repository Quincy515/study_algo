package Visitor

// Customer: 接待者
type Customer interface {
	Accept(Visitor)
}

// Visitor: 访问者
type Visitor interface {
	Visit(Customer)
}
