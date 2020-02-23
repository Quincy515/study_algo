package Chain

// 项目经理ProjectManager < 部门经理DepManager < 总经理GeneralManager
type Manager interface {
	HaveRight(money int) bool                     // 判断权限
	HandleFeeRequest(name string, money int) bool // 向后传递
}
