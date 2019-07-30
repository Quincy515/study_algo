package _4_More_about_Leetcode

func IsValid(s string) bool {
	stack := NewArrayStackNo() // 声明一个数组作为栈
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			stack.Push(rune(c)) // 压入栈顶元素
		} else {
			if stack.GetSize() == 0 { // 如果第一次压入的是右括号，没有元素进行匹配
				return false
			}

			if stack.GetSize() > 0 {
				topChar := stack.Peek() // 获取栈顶元素
				stack.Pop()             // 弹出栈顶元素
				if c == ')' && topChar != '(' {
					return false
				}
				if c == ']' && topChar != '[' {
					return false
				}
				if c == '}' && topChar != '{' {
					return false
				}
			}
		}
	}
	return stack.GetSize() == 0 // 判断栈是否为空
}
