//题目

//给定一个字符串，找出该字符串中最长回文子串的长度。

//示例 1

//输入："abc"，输出：0

//示例 2

//输入："abcbe"，输出：3

//示例 3

//输入："acdcecdcf"，输出：7

package main

func main() {
	//var s string = "abc"
	//var s string = "abcbe"
	var s string = "acdcecdcf"
	i := check(s)
	println(i)
}

func check(s string) int {
	max := 0
	if len(s) <= 3 && s[0] == s[len(s)-1] {
		return len(s)
	} else if len(s) <= 3 {
		return 0
	}
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			t := s[i:j]
			if judge(t) && max < len(t) {
				max = len(t)
			}
		}
	}
	return max
}

func judge(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
