// 题目

// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略空格、字母的大小写。

// 示例

// 输入："front man, a plan, a canal: Panama"，输出：true

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "front man, a plan, a canal: Panama"

	fmt.Println(check(s))
}

func check(s string) bool {
	s = strings.ToLower(s)
	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] > '9' || s[l] < '0' {
			if s[l] > 'z' || s[l] < 'a' {
				l++
				continue
			}
		}
		if s[r] > '9' || s[r] < '0' {
			if s[r] > 'z' || s[r] < 'a' {
				r--
				continue
			}
		}
		if s[l] != s[r] {
			return false
		} else {
			l++
			r--
		}
	}
	return true
}
