package Proxy

type Subject interface {
	Do() string // 实际业务，完成几个系统之间的挂接，业务系统检查是否欠费，检查密码是否正确
}

type RealSubject struct{}

func (RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real  RealSubject
	money int
}

func (p Proxy) Do() string {
	var res string
	if p.money > 0 {
		// 在调用真实对象之前的工作，检查缓存，判断权限，实例化真实对象等...
		res += "pre:"
		// 调用真实对象
		res += p.real.Do()
		// 调用真实对象之后的操作，如缓存结果，对结果进行处理等...
		res += ":after"
	} else {
		return "余额不足，请充值"
	}
	return res
}
