package Bridge

import "fmt"

type MessageEmail struct{}

func ViaEmail() MessageImplementer {
	return &MessageEmail{}
}

// Send: 邮件发送
func (*MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via Email", text, to)
}
