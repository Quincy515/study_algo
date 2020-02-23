package Bridge

// SMS短信
// Email

// AbstractMessage: 高度抽象比如发送速度快慢
type AbstractMessage interface {
	SendMessage(test, to string)
}

// MessageImplementer: 比如发送短信或邮件
type MessageImplementer interface {
	Send(text, to string)
}
