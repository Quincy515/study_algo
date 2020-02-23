package AbstractFactory

import "fmt"

// MySQLMainDAO: 关系型数据库mysql的OrderMainDAO实现,针对两个接口的实现
type MySQLMainDAO struct{}

func (m *MySQLMainDAO) SaveOrderMain() {
	fmt.Println("mysql main save")
}

// MySQLDetailsDAO: 关系型数据库的OrderDetailDAO实现
type MySQLDetailsDAO struct{}

func (m *MySQLDetailsDAO) SaveOrderDetails() {
	fmt.Println("mysql detail save")
}
