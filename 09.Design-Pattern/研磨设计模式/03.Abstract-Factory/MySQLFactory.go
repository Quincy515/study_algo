package AbstractFactory

// 抽象工厂实现
type MySQLFactory struct{}

func (*MySQLFactory) CreateOrderMainDAO() OrderMainDAO {
	return &MySQLMainDAO{}
}

func (*MySQLFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &MySQLDetailsDAO{}
}
