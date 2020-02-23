package SimpleFactory

// 中文: 你好
// 英文: hello

// API is interface
type API interface {
	Say(name string) string // 预留接口
}

// NewAPI return Api instance by type
func NewAPI(str string) API {
	if str == "en" {
		return &english{}
	} else if str == "zh_cn" {
		return &chinese{}
	}
	return nil
}

// 新建对象 type 新建类型别名
type chinese struct{}

// 接口的实现
func (*chinese) Say(name string) string {
	return "你好" + name
}

// English is another API implement
type english struct{}

// Say hello to name
func (*english) Say(name string) string {
	return "hello " + name
}
