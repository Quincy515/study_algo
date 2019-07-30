package _3_A_Stack_Problem_in_Leetcode

func IsValid(s string) bool {
	var stack []rune // 声明一个数组作为栈
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, rune(c)) // 压入栈顶元素
		} else {
			if len(stack) == 0 { // 如果第一次压入的是右括号，没有元素进行匹配
				return false
			}

			if len(stack) > 0 {
				topChar := stack[len(stack)-1] // 获取栈顶元素
				stack = stack[:len(stack)-1]   // 弹出栈顶元素
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
	return len(stack) == 0 // 判断栈是否为空
}
