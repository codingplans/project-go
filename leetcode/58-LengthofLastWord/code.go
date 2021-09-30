package _00_init_code

func lengthOfLastWord(s string) int {
	l := len(s)
	j, x := 0, 0
	for i := l - 1; i >= 0; i-- {
		if x == 0 && s[i] == 32 {
			j++
			println("当前是空格", s[i], x, j)
			continue
		}
		if x > 0 && s[i] == 32 {
			println("跳出判断", s[i], x, j)
			return l - i - j - 1
		}
		x++

	}

	return 0
}
