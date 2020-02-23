package Observer

type Subject struct {
	observer []Observer // 群聊
	context  string     // 更新群聊信息
}

func NewSubject() *Subject {
	return &Subject{
		observer: make([]Observer, 0),
	}
}

// Attach: 加入观察者序列
func (s *Subject) Attach(o Observer) {
	s.observer = append(s.observer, o) // 加入群聊
}

// notify: 通知，循环广播更新
func (s *Subject) notify() {
	for _, o := range s.observer { // 每个人都发一次
		o.Update(s) // 更新通知
	}
}

// UpdateContext: 更新
func (s *Subject) UpdateContext(context string) {
	s.context = context // 更新信息
	s.notify()          // 更新之后通知
}
