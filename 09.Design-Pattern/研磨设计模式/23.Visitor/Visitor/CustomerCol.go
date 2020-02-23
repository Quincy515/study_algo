package Visitor

// CustomerCol: 接待者集合
type CustomerCol struct {
	customers []Customer
}

// Accept: 追加
func (c *CustomerCol) Add(customer Customer) {
	c.customers = append(c.customers, customer)
}

// Accept: 每个服务者，都接收访问者
func (c *CustomerCol) Accept(visitor Visitor) {
	for _, customer := range c.customers {
		customer.Accept(visitor)
	}
}
