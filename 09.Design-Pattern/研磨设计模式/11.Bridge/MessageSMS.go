package Bridge

import "fmt"

type MessageSMS struct{}

func ViaSMS() MessageImplementer {
	return &MessageSMS{}
}

// Send: 短信发送
func (*MessageSMS) Send(text, to string) {
	fmt.Printf("send %s to %s via SMS", text, to)
}
