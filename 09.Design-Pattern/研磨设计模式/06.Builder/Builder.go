package Builder

// 比如启动顺序: 先打开文件->网络->数据库

type Builder interface {
	Part1() // 步骤1
	Part2() // 步骤2
	Part3() // 步骤3
}

// "1"+"2"+"3"="123"
//  1 + 2 + 3 = 6

type Director struct {
	builder Builder // 建造者接口
}

// NewDirector: 创建接口
func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

// Construct: 构造
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}
