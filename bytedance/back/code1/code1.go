//题目

//实现一个 36 进制的加法 0-9 a-z。

//示例

//输入：["abbbb","1"]，输出："abbbc"

package main

func main() {
	var s [2]string
	s[0] = "abbbb"
	s[1] = "1"
	println(transfer(s))
}

func transfer(s [2]string) string {
	s0 := len(s[0])
	s1 := len(s[1])
	var a, b int
	for i := 0; i < s0; i++ {
		if int(s[0][i]) >= 97 {
			a = a*36 + (int(s[0][i]) - 87)
		} else {
			a = a*36 + (int(s[0][i]) - 48)
		}
	}
	for j := 0; j < s1; j++ {
		if int(s[1][j]) >= 97 {
			b = b*36 + int(s[1][j]-87)
		} else {
			b = b*36 + int(s[1][j]-48)
		}
	}
	c := a + b
	var res string
	for c > 0 {
		t := c % 36
		if t > 9 {
			res += string(t + 87)
		} else {
			res += string(t + 48)
		}
		c /= 36
	}
	var r string
	for _, ch := range res {
		r = string(ch) + r
	}
	return r
}
