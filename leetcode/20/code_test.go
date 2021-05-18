package _00_init_code

import (
	"fmt"
	"testing"
)

type test struct {
	IntEs []int
	Lists []int
	K     int
	S     string
}

var tests = []test{
	{S: "()[]{}"},
	{S: "([]}"},
	{S: "([)]"},
	{S: "({}{()})"},
	{S: "([}}])"},
	{S: "()[]{{}}[]()"},
}

// 20. 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
// 示例 1：
//
// 输入：s = "()"
// 输出：true
// 示例 2：
//
// 输入：s = "()[]{}"
// 输出：true
// 示例 3：
//
// 输入：s = "(]"
// 输出：false
// 示例 4：
//
// 输入：s = "([)]"
// 输出：false
// 示例 5：
//
// 输入：s = "{[]}"
// 输出：true
//
//
// 提示：
//
// 1 <= s.length <= 104
// s 仅由括号 '()[]{}' 组成

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化")
		// fmt.Println('[', ']')
		// fmt.Println('{', '}')
		// fmt.Println('(', ')')
		pre := isValid(tests[k1].S)
		fmt.Println("结果：", pre)

	}
}

// 思路：像回文链一样，每次入栈 ，然后和栈顶作对比，相等则消除，不等则入栈，最后判断 栈长度是否为 0
func isValid(s string) bool {
	md := make(map[byte]byte, 3)
	md['['] = 93
	md['{'] = 125
	md['('] = 41

	stack := []byte{}
	flag := 0
	for v := range s {
		l := len(stack)
		if l > 0 {
			if stack[l-1] == s[v] && flag == 0 {
				stack = stack[:l-1]
				continue
			}
		}
		if v1, ok := md[s[v]]; ok {
			stack = append(stack, v1)
			flag = 0
		} else {
			flag = 1
			stack = append(stack, s[v])
		}

	}
	return len(stack) == 0
}
