package Mediator

// Mediator: 单例模式，通过一个类控制所有接口
type Mediator struct {
	CD    *CDDriver
	CPU   *CPU
	Video *VideoCard
	Sound *SoundCard
}

var mediator *Mediator

// GetMediatorInstance: 单例模式
func GetMediatorInstance() *Mediator {
	if mediator == nil { // 没有构造
		mediator = &Mediator{} // 构造初始化
	}
	return mediator
}

func (m *Mediator) changed(i interface{}) {
	switch instance := i.(type) {
	case *CDDriver:
		m.CPU.Process(instance.Data)
	case *CPU:
		m.Sound.Play(instance.Sound)
		m.Video.Display(instance.Video)
	}
}
