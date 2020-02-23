package AbstractFactory

import "fmt"

// mysql针对两个接口的实现
type OracleMainDAO struct{}

func (*OracleMainDAO) SaveOrderMain() {
	fmt.Println("oracle main save")
}

type OracleDetailsDAO struct{}

func (*OracleDetailsDAO) SaveOrderDetails() {
	fmt.Println("oracle detail save")
}
