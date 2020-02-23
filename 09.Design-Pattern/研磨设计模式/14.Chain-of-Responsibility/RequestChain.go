package Chain

type RequestChain struct {
	Manager
	successor *RequestChain // 处理成功标记
}

// SetSuccessor: 判断责任链在哪里结束
func (r *RequestChain) SetSuccessor(endrc *RequestChain) {
	r.successor = endrc
}

// HaveRight: 判断权限
func (r *RequestChain) HaveRight(money int) bool {
	return true
}

// HandleFeeRequest: 向后传递，有权利当前处理，没有权利放弃，向后传递
func (r *RequestChain) HandleFeeRequest(name string, money int) bool {
	if r.Manager.HaveRight(money) {
		return r.Manager.HandleFeeRequest(name, money)
	}
	if r.successor != nil {
		return r.successor.HandleFeeRequest(name, money)
	}
	return false
}
