package _00_init_code

import (
	"fmt"
	"testing"
)

type test struct {
	IntEs []int
	Lists []int
	Str   string
	Str2  string
	K     int
}

var tests = []test{
	{
		Str:  "horse",
		Str2: "ros",
	},
}

// 2021 0521 未完成！！！

// 72. 编辑距离
// 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
//
// 你可以对一个单词进行如下三种操作：
//
// 插入一个字符
// 删除一个字符
// 替换一个字符
//
//
// 示例 1：
//
// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
// 示例 2：
//
// 输入：word1 = "inten tion", word2 = "execu tion"
// 输出：5
// 解释：
// intention -> inention (删除 't')
// inention -> enention (将 'i' 替换为 'e')
// enention -> exention (将 'n' 替换为 'x')
// exention -> exection (将 'n' 替换为 'c')
// exection -> execution (插入 'u')
//
//
// 提示：
//
// 0 <= word1.length, word2.length <= 500
// word1 和 word2 由小写英文字母组成

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化")

		minDistance(tests[k1].Str, tests[k1].Str2)
		pre := 1
		fmt.Println("结果：", pre)

	}
}

func minDistance(word1 string, word2 string) int {
	fmt.Println(word1, word2)

	l2 := len(word2)
	for i := 0; i < l2; i++ {
		if word1[i] == word2[i] {
			continue
		}
		// w3:=append(word1[:i-1],word2[i])
		// aa := word1[:i-1] + word2[i]
		// aa := word1[:i-1] + word2[i]

		m := i - 1
		if m < 0 {
			m = 0
		}
		fmt.Println(word1[:m], 222, word1[i+1:])
	}

	return 0

}

func insert(word1, k string, index int) {

}

func del(word1, k string, index int) {

}

func upData(word1, k, k2 string, index int) {

}
