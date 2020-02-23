package AbstractFactory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	var factory DAOFactor
	factory = &MySQLFactory{}
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetails()
}

func TestAbstractFactoryOracle(t *testing.T) {
	var factory DAOFactor
	factory = &OracleFactory{} // different
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetails()
}
