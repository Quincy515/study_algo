package AbstractFactory

// 抽象工厂
type OracleFactory struct{}

func (*OracleFactory) CreateOrderMainDAO() OrderMainDAO {
	return &OracleMainDAO{}
}

func (*OracleFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &OracleDetailsDAO{}
}
