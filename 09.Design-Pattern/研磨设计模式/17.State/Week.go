package State

// Week: 每个星期的行为
type Week interface {
	Today()
	Next(*DayContext)
}

// DayContext: 星期的数据结构
type DayContext struct {
	today Week
}

func NewDayContext() *DayContext {
	return &DayContext{
		today: &Sunday{},
	}
}

// Today: 今天
func (d *DayContext) Today() {
	d.today.Today()
}

// Next: 明天
func (d *DayContext) Next(*DayContext) {
	d.today.Next(d)
}
