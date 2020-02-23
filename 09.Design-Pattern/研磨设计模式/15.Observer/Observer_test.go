package Observer

func ExampleObserver() {
	subject := NewSubject()         // 创建群聊
	reader1 := NewReader("reader1") // 创建订阅者
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")

	subject.Attach(reader1) // 加入群聊
	subject.Attach(reader2)
	subject.Attach(reader3)

	subject.UpdateContext("observer mode") // 更新信息进行通知
	// Output:
	// reader1 receive observer mode
	// reader2 receive observer mode
	// reader3 receive observer mode
}
