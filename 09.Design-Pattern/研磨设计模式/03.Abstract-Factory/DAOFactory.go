package AbstractFactory

// 数据库连接
// sqlserver、mysql、oracle
// OrderMainDAO: 订单主记录
type OrderMainDAO interface {
	SaveOrderMain() // 保存订单
	//DeleteOrderMain() // 删除订单
	//SearchOrderMain() // 搜索订单
}

// OrderDetailDAO: 订单详情记录
type OrderDetailDAO interface {
	SaveOrderDetails() // 保存
}

// DAOFactor: DAO抽象模式工厂接口 完全抽象的接口
type DAOFactor interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}
